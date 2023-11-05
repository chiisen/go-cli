package cmd
 
import (
   "fmt"
   "github.com/spf13/cobra"
   "os"
)
 
var name string
 
var rootCmd = &cobra.Command{
   Use: "hello",
   Run: func(cmd *cobra.Command, args []string) {
      fmt.Println("hello ", name)
   },
}
 
func Execute() {
   err := rootCmd.Execute()
   if err != nil {
      os.Exit(1)
   }
}
 
func init() {
   rootCmd.Flags().StringVarP(&name, "name", "", "world", "")
}
