//  Paxos 算法就是通过两个阶段确定一个决议：

// Phase1：确定谁的编号最高，只有编号最高者才有权利提交 Proposal（提议：给定的具体值）
// Phase2：编号最高者提交 Proposal，如果没有其他节点提出更高编号的 Proposal，则该提案会被顺利通过，否则整个过程就会重来。

// Phase1

func (px *Paxos)Prepare(args *PrepareArgs, reply *PrepareReply) error {
	px.mu.Lock()
	defer px.mu.Unlock()
	// 通过 RPC 调用，询问每一台机器，当前的这个提议能不能通过
	round, exist := px.rounds[args.Seq]
	if !exist {
		// new seq of commit, so need new
		px.rounds[args.Seq] = px.newInstance()
		round, _ := px.rounds[args.Seq]
		reply.Err = OK
	} else {
		// 当前提交的编号大于之前的其他机器 Prepare 的编号
		if args.PNum > round.proposeNumber {
			reply.Err = OK
		} else {
			reply.Err = Reject
		}
	}
	
	// 如果本次提议是通过的，那么还需返回给提议者，已经通过提议和确定的值。
	if reply.Err == OK {
		reply.AcceptPnum = round.acceptorNumber
		reply.AcceptValue = round.acceptValue
		px.rounds[args.Seq].proposeNumber = args.PNum
	} else {
		// reject
	}
	return nil
}

// Phase2
func (px Paxos)Accept(args *AcceptArgs, reply *AcceptReply) error {
	px.mu.Lock()
	defer px.mu.Unlock()
	round, exist := px.rounds[args.Seq]
	if !exist {
		px.rounds[args.Seq] = px.newInstance()
		reply.Err = OK
	} else {
		if args.PNum >= round.proposeNumber {
			reply.Err = OK
		} else {
			reply.Err = Reject
		}
	}

	// 如果提议通过，那么就需设置当轮的提议编号和提议的值。
	if reply.Err == OK {
		px.rounds[args.Seq].acceptorNumber = args.PNum
		px.rounds[args.Seq].proposeNumber = args.PNum
		px.rounds[args.Seq].acceptValue = args.Value
	} else {
		// reject
	}
	return nil
}

rounds map[int]*Round // cacehe each round paxos result key is seq value is value
completes []int // maintain peer min seq of comleted

// Decide 方法，用于提议者来确定某个值，这个映射到分布式里面的状态机的应用。
func (px *Paxos)Decide(args *DecideArgs, reply *DecideReply) error {
	px.mu.Lock()
	defer px.mu.Unlock()
	_, exist := px.rounds[args.Seq]
	if !exist {
		px.rounds[args.Seq] = px.newInstance()
	}
	px.rounds[args.Seq].acceptorNumber = args.PNum
	px.rounds[args.Seq].acceptValue = args.Value
	px.rounds[args.Seq].proposeNumber = args.PNum
	px.rounds[args.Seq].state = Decided
	px.completes[args.Me] = args.Done
	return nil
}