package youtubedownloader



var randomCmd = &cobra.Command{
    Use:   "random",
    Short: "Get a random dad joke",
    Long:  `This command fetches a random dad joke from the icanhazdadjoke api`,
    Run: func(cmd *cobra.Command, args []string) {
        ...
    },
}
