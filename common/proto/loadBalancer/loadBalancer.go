package loadBalancer

func RequestToProto(addr string) *Request {
	return &Request{
		Addr: addr,
	}
}
