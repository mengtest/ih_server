package main

import (
	"ih_server/libs/log"
	"ih_server/src/table_config"

	"sync"
	"time"

	"ih_server/proto/gen_go/client_message"
	"ih_server/proto/gen_go/client_message_id"
	"ih_server/proto/gen_go/rpc_message"
	"ih_server/src/rpc_proto"

	"github.com/golang/protobuf/proto"
)

func GetCarnivalCurrRoundAndRemainSeconds() (round, remain_seconds int32) {
	now_time := time.Now()
	for i := 0; i < len(carnival_table_mgr.Array); i++ {
		c := carnival_table_mgr.Array[i]
		if c != nil {
			if int32(now_time.Unix()) >= c.StartTime && int32(now_time.Unix()) <= c.EndTime {
				round = c.Round
				remain_seconds = c.EndTime - int32(now_time.Unix())
				break
			}
		}
	}
	return
}

func (this *Player) _carnival_data_check(send_notify bool) (round, remain_seconds int32) {
	round, remain_seconds = GetCarnivalCurrRoundAndRemainSeconds()
	prev_round := dbc.Carnival.GetRow().GetRound()

	if round > prev_round {
		round_reset_tasks := carnival_task_table_mgr.GetRoundResetTasks()
		if round_reset_tasks != nil {
			for _, t := range round_reset_tasks {
				if this.db.Carnivals.HasIndex(t.Id) {
					this.db.Carnivals.SetValue(t.Id, 0)
					if send_notify {
						this.carnival_task_data_notify(t.Id, 0, 0)
					}
				}
			}
		}
		dbc.Carnival.GetRow().SetRound(round)
	}

	now_time := time.Now()
	last_day_reset_time := this.db.CarnivalCommon.GetDayResetTime()
	if last_day_reset_time == 0 {
		last_day_reset_time = int32(now_time.Unix())
		this.db.CarnivalCommon.SetDayResetTime(int32(now_time.Unix()))
	}

	last_unix := time.Unix(int64(last_day_reset_time), 0)
	if last_unix.Year() < now_time.Year() || last_unix.Month() < now_time.Month() || last_unix.Day() < now_time.Day() {
		day_reset_tasks := carnival_task_table_mgr.GetDayResetTasks()
		if day_reset_tasks != nil {
			for _, t := range day_reset_tasks {
				if this.db.Carnivals.HasIndex(t.Id) {
					this.db.Carnivals.SetValue(t.Id, 0)
					if send_notify {
						this.carnival_task_data_notify(t.Id, 0, 0)
					}
				}
			}
		}
		this.db.CarnivalCommon.SetDayResetTime(int32(now_time.Unix()))
	}

	return
}

func (this *Player) carnival_data() int32 {
	round, remain_seconds := this._carnival_data_check(false)

	tasks := carnival_task_table_mgr.Array
	var task_list []*msg_client_message.CarnivalTaskData
	if tasks != nil {
		for _, t := range tasks {
			value, o := this.db.Carnivals.GetValue(t.Id)
			value2, o := this.db.Carnivals.GetValue2(t.Id)
			if !o {
				value = 0
				value2 = 0
			}
			task_list = append(task_list, &msg_client_message.CarnivalTaskData{
				Id:     t.Id,
				Value:  value,
				Value2: value2,
			})
		}
	}

	response := &msg_client_message.S2CCarnivalDataResponse{
		Round:         round,
		RemainSeconds: remain_seconds,
		TaskList:      task_list,
		InviteCode:    invite_code_generator.Generate(this.Id),
	}
	this.Send(uint16(msg_client_message_id.MSGID_S2C_CARNIVAL_DATA_RESPONSE), response)

	log.Trace("Player %v carnival data %v", this.Id, response)

	return 1
}

func (this *Player) carnival_task_data_notify(id, value, value2 int32) {
	this.Send(uint16(msg_client_message_id.MSGID_S2C_CARNIVAL_TASK_DATA_NOTIFY), &msg_client_message.S2CCarnivalTaskDataNotify{
		Data: &msg_client_message.CarnivalTaskData{
			Id:     id,
			Value:  value,
			Value2: value2,
		},
	})
}

func (this *Player) carnival_task_is_finished(task *table_config.XmlCarnivalTaskItem) bool {
	value, o := this.db.Carnivals.GetValue(task.Id)
	if !o || value < task.EventCount {
		return false
	}
	return true
}

func (this *Player) carnival_task_do_once(task *table_config.XmlCarnivalTaskItem) (value, value2 int32) {
	if !this.db.Carnivals.HasIndex(task.Id) {
		this.db.Carnivals.Add(&dbPlayerCarnivalData{
			Id: task.Id,
		})
	}

	var reward bool
	if task.EventType == table_config.CARNIVAL_EVENT_INVITE {
		value2 = this.db.Carnivals.IncbyValue2(task.Id, 1)
		if value2 >= task.Param1 {
			value = this.db.Carnivals.IncbyValue(task.Id, 1)
			this.db.Carnivals.SetValue2(task.Id, 0)
			reward = true
		}
	} else {
		value = this.db.Carnivals.IncbyValue(task.Id, 1)
		reward = true
	}

	if reward {
		if task.RewardMailId > 0 {
			RealSendMail(nil, this.Id, MAIL_TYPE_SYSTEM, task.RewardMailId, "", "", task.Reward, 0)
		} else {
			this.add_resources(task.Reward)
		}
	}

	return
}

func (this *Player) carnival_task_set(id int32) int32 {
	round, _ := this._carnival_data_check(true)
	if round <= 0 {
		log.Error("no carnival task doing in this time")
		return int32(msg_client_message.E_ERR_CARNIVAL_NOT_DOING)
	}

	task := carnival_task_table_mgr.Get(id)
	if task == nil {
		return int32(msg_client_message.E_ERR_CARNIVAL_TASK_NOT_FOUND)
	}
	if !(task.EventType == table_config.CARNIVAL_EVENT_COMMENT || task.EventType == table_config.CARNIVAL_EVENT_FOCUS_COMMUNITY || task.EventType == table_config.CARNIVAL_EVENT_SHARE) {
		log.Error("carnival task %v with event type %v cant set progress", id, task.EventType)
		return int32(msg_client_message.E_ERR_CARNIVAL_TASK_CANT_SET)
	}

	if this.carnival_task_is_finished(task) {
		log.Error("Player %v already complete carnival task %v", this.Id, id)
		return int32(msg_client_message.E_ERR_CARNIVAL_TASK_ALREADY_FINISHED)
	}

	value, value2 := this.carnival_task_do_once(task)

	this.Send(uint16(msg_client_message_id.MSGID_S2C_CARNIVAL_TASK_SET_RESPONSE), &msg_client_message.S2CCarnivalTaskSetResponse{
		TaskId: id,
	})

	this.carnival_task_data_notify(id, value, value2)

	log.Trace("Player %v carnival task %v progress %v/%v", this.Id, id, value, task.EventCount)

	return 1
}

func (this *Player) carnival_item_exchange(task_id int32) int32 {
	round, _ := this._carnival_data_check(true)
	if round <= 0 {
		log.Error("no carnival task doing in this time")
		return int32(msg_client_message.E_ERR_CARNIVAL_NOT_DOING)
	}

	task := carnival_task_table_mgr.Get(task_id)
	if task == nil {
		return int32(msg_client_message.E_ERR_CARNIVAL_TASK_NOT_FOUND)
	}

	if this.carnival_task_is_finished(task) {
		log.Error("Player %v already complete carnival task %v", this.Id, task_id)
		return int32(msg_client_message.E_ERR_CARNIVAL_TASK_ALREADY_FINISHED)
	}

	var items = []int32{task.Param1, task.Param2, task.Param3, task.Param4}
	if !this.check_resources(items) {
		log.Error("Player %v item exchange not enough for carnival task %v", this.Id, task_id)
		return int32(msg_client_message.E_ERR_PLAYER_ITEM_NUM_NOT_ENOUGH)
	}

	this.cost_resources(items)
	value, value2 := this.carnival_task_do_once(task)

	this.Send(uint16(msg_client_message_id.MSGID_S2C_CARNIVAL_ITEM_EXCHANGE_RESPONSE), &msg_client_message.S2CCarnivalItemExchangeResponse{
		TaskId: task_id,
	})

	this.carnival_task_data_notify(task_id, value, value2)

	log.Trace("Player %v item exchanged for carnival task %v progress %v/%v", this.Id, task_id, value, task.EventCount)

	return 1
}

func carnival_get_task_by_type(task_type int32) *table_config.XmlCarnivalTaskItem {
	var task *table_config.XmlCarnivalTaskItem
	task_array := carnival_task_table_mgr.Array
	for i := 0; i < len(task_array); i++ {
		if task_array[i] == nil {
			continue
		}
		if task_array[i].EventType == task_type {
			task = task_array[i]
			break
		}
	}
	return task
}

type InviteCodeGenerator struct {
	Source     []byte
	Char2Index map[byte]int
	tmp_index  []int
	tmp_length int
	locker     sync.RWMutex
}

var invite_code_generator InviteCodeGenerator

func (this *InviteCodeGenerator) Init() {
	this.Source = []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	this.Char2Index = make(map[byte]int)
	source_bytes := []byte(this.Source)
	for i := 0; i < len(source_bytes); i++ {
		this.Char2Index[source_bytes[i]] = i
	}
	this.tmp_index = make([]int, 10)
}

func (this *InviteCodeGenerator) Generate(id int32) (code string) {
	l := len(this.Source)
	if l == 0 {
		log.Error("InviteCodeGenerator not init")
		return ""
	}

	this.locker.Lock()
	defer this.locker.Unlock()

	a := int(id)
	for {
		n := 0
		t := a
		for t >= l {
			t /= l
			n += 1
		}
		if this.tmp_length == 0 {
			this.tmp_length = n + 1
			log.Trace("@@@@@ InviteCodeGenerator   this.tmp_length %v, n %v", this.tmp_length, n)
		}
		log.Trace("@@@@ InviteCodeGenerator   n = %v   t = %v", n, t)
		this.tmp_index[this.tmp_length-1-n] = t
		if n > 0 {
			a -= (t * power_n(l, n))
		}
		log.Trace("@@@@ InviteCodeGenerator   a = %v", a)

		if a < l {
			this.tmp_index[this.tmp_length-1] = a
			break
		}
	}

	for i := 0; i < this.tmp_length; i++ {
		code += string(this.Source[this.tmp_index[i]])
	}

	for i := 0; i < this.tmp_length; i++ {
		this.tmp_index[i] = 0
	}
	this.tmp_length = 0

	return
}

func (this *InviteCodeGenerator) GetId(code string) (id int32) {
	source_length := len(this.Source)
	code_bytes := []byte(code)
	for i := 0; i < len(code_bytes); i++ {
		idx, o := this.Char2Index[code_bytes[i]]
		if !o {
			log.Error("InviteCodeGenerator code include byte %v is invalid", code_bytes[i])
			return 0
		}
		id += int32((idx) * (power_n(source_length, (len(code_bytes) - i - 1))))
	}
	return
}

func (this *Player) carnival_share() int32 {
	round, _ := this._carnival_data_check(true)
	if round <= 0 {
		log.Error("no carnival task doing in this time")
		return int32(msg_client_message.E_ERR_CARNIVAL_NOT_DOING)
	}

	// 分享任务
	task := carnival_get_task_by_type(table_config.CARNIVAL_EVENT_SHARE)
	if task == nil {
		log.Error("Not found carnival share task")
		return int32(msg_client_message.E_ERR_CARNIVAL_TASK_NOT_FOUND)
	}

	if this.carnival_task_is_finished(task) {
		log.Error("Player %v carnival task %v already finished", this.Id, task.Id)
		return int32(msg_client_message.E_ERR_CARNIVAL_TASK_ALREADY_FINISHED)
	}

	// 生成邀请码
	invite_code := invite_code_generator.Generate(this.Id)
	value, value2 := this.carnival_task_do_once(task)

	this.Send(uint16(msg_client_message_id.MSGID_S2C_CARNIVAL_SHARE_RESPONSE), &msg_client_message.S2CCarnivalShareResponse{
		InviteCode: invite_code,
	})
	this.carnival_task_data_notify(task.Id, value, value2)

	log.Trace("Player %v carnival share task %v progress %v/%v, invite code %v", this.Id, task.Id, value, task.EventCount, invite_code)

	return 1
}

func (this *Player) carnival_invite_tasks_check() bool {
	invite_tasks := carnival_task_table_mgr.GetInviteTasks()
	if invite_tasks == nil {
		return false
	}

	var do bool
	for _, t := range invite_tasks {
		if this.carnival_task_is_finished(t) {
			continue
		}
		value, value2 := this.carnival_task_do_once(t)
		this.carnival_task_data_notify(t.Id, value, value2)
		do = true
		log.Trace("Player %v carnival invite task %v progress %v/%v", this.Id, t.Id, value, t.EventCount)
	}

	return do
}

func (this *Player) carnival_be_invited(invite_code string) int32 {
	round, _ := this._carnival_data_check(true)
	if round <= 0 {
		log.Error("no carnival task doing in this time")
		return int32(msg_client_message.E_ERR_CARNIVAL_NOT_DOING)
	}

	task := carnival_get_task_by_type(table_config.CARNIVAL_EVENT_BE_INVITED_REWARD)
	if task == nil {
		log.Error("Not found carnival be invited task")
		return int32(msg_client_message.E_ERR_CARNIVAL_TASK_NOT_FOUND)
	}

	if this.db.InviteCodess.HasIndex(invite_code) {
		log.Error("Player %v already used invite code %v", this.Id, invite_code)
		return int32(msg_client_message.E_ERR_CARNIVAL_TASK_ALREADY_FINISHED)
	}

	if this.carnival_task_is_finished(task) {
		log.Error("Player %v carnival task %v already finished", this.Id, task.Id)
		return int32(msg_client_message.E_ERR_CARNIVAL_TASK_ALREADY_FINISHED)
	} else {
		is_invited, err_code := remote_carnival_get_is_invited(this.Id, this.UniqueId, invite_code)
		if is_invited {
			log.Error("Player %v carnival remote invited with task %v already finished", this.Id, task.Id)
			return int32(msg_client_message.E_ERR_CARNIVAL_TASK_ALREADY_FINISHED)
		}
		if err_code < 0 {
			return err_code
		}
	}

	// 邀请者任务检测
	inviter_id := invite_code_generator.GetId(invite_code)
	if inviter_id <= 0 {
		log.Error("Player %v provide invite code %v invalid", this.Id, invite_code)
		return int32(msg_client_message.E_ERR_CARNIVAL_TASK_INVITE_CODE_INVALID)
	}

	if inviter_id == this.Id {
		log.Error("Player %v cant invite self on carnival", this.Id)
		return -1
	}

	inviter := player_mgr.GetPlayerById(inviter_id)
	if inviter == nil {
		_, err_code := remote_carnival_be_invited(this.Id, inviter_id)
		if err_code < 0 {
			log.Error("Player %v remote carnival invite by %v err %v", this.Id, inviter_id, err_code)
			return err_code
		}
	} else {
		if !inviter.carnival_invite_tasks_check() {
			log.Error("Player %v use the invite code %v deprecated", this.Id, invite_code)
			return int32(msg_client_message.E_ERR_CARNIVAL_TASK_INVITE_CODE_DEPRECATED)
		}
	}

	value, value2 := this.carnival_task_do_once(task)
	this.db.InviteCodess.Add(&dbPlayerInviteCodesData{
		Code: invite_code,
	})

	this.Send(uint16(msg_client_message_id.MSGID_S2C_CARNIVAL_BE_INVITED_RESPONSE), &msg_client_message.S2CCarnivalBeInvitedResponse{
		InviteCode: invite_code,
	})
	this.carnival_task_data_notify(task.Id, value, value2)

	log.Trace("Player %v carnival be invite task %v use invite code %v, progress %v/%v", this.Id, task.Id, invite_code, value, task.EventCount)

	return 1
}

func C2SCarnivalDataHandler(p *Player, msg_data []byte) int32 {
	var req msg_client_message.C2SCarnivalDataRequest
	err := proto.Unmarshal(msg_data, &req)
	if err != nil {
		log.Error("Unmarshal msg failed err(%s)!", err.Error())
		return -1
	}
	return p.carnival_data()
}

func C2SCarnivalTaskSetHandler(p *Player, msg_data []byte) int32 {
	var req msg_client_message.C2SCarnivalTaskSetRequest
	err := proto.Unmarshal(msg_data, &req)
	if err != nil {
		log.Error("Unmarshal msg failed err(%s)!", err.Error())
		return -1
	}
	return p.carnival_task_set(req.GetTaskId())
}

func C2SCarnivalItemExchangeHandler(p *Player, msg_data []byte) int32 {
	var req msg_client_message.C2SCarnivalItemExchangeRequest
	err := proto.Unmarshal(msg_data, &req)
	if err != nil {
		log.Error("Unmarshal msg failed err(%s)!", err.Error())
		return -1
	}
	return p.carnival_item_exchange(req.GetTaskId())
}

func C2SCarnivalShareHandler(p *Player, msg_data []byte) int32 {
	var req msg_client_message.C2SCarnivalShareRequest
	err := proto.Unmarshal(msg_data, &req)
	if err != nil {
		log.Error("Unmarshal msg failed err(%s)!", err.Error())
		return -1
	}
	return p.carnival_share()
}

func C2SCarnivalBeInvitedHander(p *Player, msg_data []byte) int32 {
	var req msg_client_message.C2SCarnivalBeInvitedRequest
	err := proto.Unmarshal(msg_data, &req)
	if err != nil {
		log.Error("Unmarshal msg failed err(%s)!", err.Error())
		return -1
	}
	return p.carnival_be_invited(req.GetInviteCode())
}

// ---------------------------------- remote ----------------------------------
// 嘉年华是否被邀请
func remote_carnival_get_is_invited(from_player_id int32, unique_id, invite_code string) (is_invited bool, err_code int32) {
	var req = msg_rpc_message.G2GCarnivalIsInvitedRequest{
		PlayerUniqueId: unique_id,
		InviteCode:     invite_code,
	}

	req_data, err := _marshal_msg(&req)
	if err != nil {
		err_code = -1
		return
	}

	datas := rpc_proto.RpcBroadcastGet(hall_server.rpc_client, "G2G_CommonProc.BroadcastGet", from_player_id, int32(msg_rpc_message.MSGID_G2G_CARNIVAL_IS_INVITED_REQUEST), req_data)
	if datas == nil || len(datas) == 0 {
		log.Error("remote carnival is invited get result empty")
		err_code = -1
		return
	}

	var response msg_rpc_message.G2GCarnivalIsInvitedResponse
	for i := 0; i < len(datas); i++ {
		if datas[i].ErrorCode < 0 {
			err_code = datas[i].ErrorCode
			log.Error("remote carnival is invited error %v with index %v", datas[i].ErrorCode, i)
			return
		}
		err = _unmarshal_msg(datas[i].ResultData, &response)
		if err != nil {
			err_code = -1
			return
		}
		if response.IsInvited {
			is_invited = true
		}
	}

	err_code = 1
	return
}

// 嘉年华是否被邀请返回
func remote_carnival_get_is_invited_response(from_player_id int32, req_data []byte) (resp_data []byte, err_code int32) {
	var req msg_rpc_message.G2GCarnivalIsInvitedRequest
	err := _unmarshal_msg(req_data, &req)
	if err != nil {
		err_code = -1
		return
	}

	var is_invited bool
	p := player_mgr.GetPlayerByUid(req.GetPlayerUniqueId())
	if p != nil {
		if p.db.InviteCodess.HasIndex(req.GetInviteCode()) {
			is_invited = true
		}
	}

	if err_code >= 0 {
		var response = msg_rpc_message.G2GCarnivalIsInvitedResponse{
			IsInvited: is_invited,
		}
		resp_data, err = _marshal_msg(&response)
		if err != nil {
			err_code = -1
			return
		}
	}

	err_code = 1
	return
}

// 嘉年華被邀請
func remote_carnival_be_invited(from_player_id, to_player_id int32) (resp *msg_rpc_message.G2GCarnivalBeInvitedResponse, err_code int32) {
	var req msg_rpc_message.G2GCarnivalBeInvitedRequest
	var response msg_rpc_message.G2GCarnivalBeInvitedResponse
	err_code = RemoteGetUsePB(from_player_id, rpc_proto.OBJECT_TYPE_PLAYER, to_player_id, int32(msg_rpc_message.MSGID_G2G_CARNIVAL_BE_INVITED_REQUEST), &req, &response)
	resp = &response
	return
}

// 嘉年華被邀請返回
func remote_carnival_be_invited_response(to_player_id int32, req_data []byte) (resp_data []byte, err_code int32) {
	var req msg_rpc_message.G2GCarnivalBeInvitedRequest
	err := _unmarshal_msg(req_data, &req)
	if err != nil {
		err_code = -1
		return
	}

	player := player_mgr.GetPlayerById(to_player_id)
	if player == nil {
		err_code = int32(msg_client_message.E_ERR_PLAYER_NOT_EXIST)
		log.Error("remote request carnival be invited by id %v not found", to_player_id)
		return
	}

	if !player.carnival_invite_tasks_check() {
		err_code = int32(msg_client_message.E_ERR_CARNIVAL_TASK_INVITE_CODE_DEPRECATED)
		log.Error("remote carnival invite tasks check failed")
		return
	}

	var response = msg_rpc_message.G2GCarnivalBeInvitedResponse{
		InviteCode: req.GetInviteCode(),
	}

	resp_data, err = _marshal_msg(&response)
	if err != nil {
		err_code = -1
		return
	}

	err_code = 1
	return
}
