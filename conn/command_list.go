package conn

const listArgSelector int = 1

func cmdList(args commandArgs, c *Conn) {
	if !c.assertAuthenticated(args.ID()) {
		return
	}

	if args.Arg(listArgSelector) == "" {
		// Blank selector means request directory separator
		c.writeResponse("", "LIST (\\Noselect) \"/\" \"\"")
	} else if args.Arg(listArgSelector) == "*" {
		// List all mailboxes requested
		for _, mailbox := range c.User.Mailboxes() {
			c.writeResponse("", "LIST () \"/\" \""+mailbox.Name()+"\"")
		}
	}
	c.writeResponse(args.ID(), "OK LIST completed")
}
