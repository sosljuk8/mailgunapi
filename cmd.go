package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"mailgunapi/client"
)

// RootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "home",
	Short: "A brief description of your application",
	Long:  `A longer description that spans multiple lines`,
}

var fetcherCmd = &cobra.Command{
	Use:   "fetcher",
	Short: "Fetch data from remote sources",
	Long:  `This command fetches data from remote sources and stores it in the database`,
	Run: func(cmd *cobra.Command, args []string) {

		mgc := client.NewMgClient()

		req, err := mgc.FormatEmailRequest("Excited User mailgun@sandbox6c87a23cb52845a989d24801f83c6ecb.mailgun.org",
			"ddd@sandbox6c87a23cb52845a989d24801f83c6ecb.mailgun.org", "Test email", "This is a test email!")
		if err != nil {
			panic(err)
		}

		res, err := mgc.Client.Do(req)
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}

		fmt.Println(string(body))
	},
}

var mailerCmd = &cobra.Command{
	Use:   "mailer",
	Short: "Run mailer sends",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mailer works, recipients: ", args)
		//// create database ORM client
		//c := orm.NewClient()
		//defer c.Close()
		//
		//// create high-level database store
		//s := orm.NewStore(c)
		//
		//// initialize jobs for crawler
		//jobs := parse.NewJobs()
		//jobs.Add(plugin.NewHydacJob())
		//
		//// create crawler instance
		//p := parse.NewCrawler(s, jobs)
		//
		//// run crawler until stopped
		//p.Run()
	},
}
