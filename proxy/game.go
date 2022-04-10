package proxy

type Game interface {
	PlayerMaxNum() int
	PlayerJoin(userKey interface{}, view bool)
	PlayerOp(userKey interface{}, view bool)
	PlayerExit(userKey interface{}, view bool)
	View() bool
	DisbandTs() int64
	IpLimit() bool
	DistanceLimit() int
	Finish(disband bool)
}
