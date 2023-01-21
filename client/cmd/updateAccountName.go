package cmd

var updateAccountName = &cobra.Command{
	Use: "updateAccountName",
	Short: "Update the account name of an existing account",
	Args: cobra.ExactArgs(4),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("updateAccountName called")
		flag.Parse()

		conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		defer conn.Close()
		c := server.NewMsgClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		r, err := c.UpdateAccountName(ctx, &server.MsgUpdateAccountName{Porfolio: args[0], Name: args[1], NewName: args[2]})
		if err != nil {
			fmt.Printf("Could not update portfolio: %v\n", err)
			return
		}
		fmt.Println("Updated Portfolio:")
		r.GetPortfolio().Println()
	}
}