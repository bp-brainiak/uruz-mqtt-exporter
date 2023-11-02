/*
SPDX-License-Identifier: GPL-2.0-only
Copyright © 2023 Ulises Ruz Puga ulises.ruz@gmail.com

This program is free software; you can redistribute it and/or
modify it under the terms of the GNU General Public License
as published by the Free Software Foundation; either version 2
of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"github.com/bp-brainiak/uruz-mqtt-exporter/mqttlogic"
	"log"

	"github.com/spf13/cobra"
)

var server, user, password string
var port int64
var topics []string
var verbose bool
var process_error error

// connectCmd represents the connect command
var connectCmd = &cobra.Command{
	Use:   "connect -s [server] -u [user] -p [password] -t [topics..]  ",
	Short: "connect to the mqtt broker and subscribe to the given topics",
	Long:  `Establish the connection to the mqtt server and subscribe to the topics`,
	Run: func(cmd *cobra.Command, args []string) {
		server, process_error = cmd.Flags().GetString("server")
		if process_error != nil {
			log.Panic("Error processing server parameter", process_error)
		}
		port, process_error = cmd.Flags().GetInt64("port")
		if process_error != nil {
			log.Panic("Error processing port parameter", process_error)
		}
		user, process_error = cmd.Flags().GetString("user")
		if process_error != nil {
			log.Panic("Error processing user parameter", process_error)
		}
		password, process_error = cmd.Flags().GetString("password")
		if process_error != nil {
			log.Panic("Error processing pass parameter", process_error)
		}
		topics, process_error = cmd.Flags().GetStringSlice("topic")
		if process_error != nil {
			log.Panic("Error processing topic parameter", process_error)
		}
		mqConfig := new(mqttlogic.MqttData)
		mqConfig.Server = server
		mqConfig.Port = port
		mqConfig.User = user
		mqConfig.Pass = password
		mqConfig.Topics = topics
		mqConfig.Verbose = verbose
		mqttlogic.SetConfigData(*mqConfig)
		mqttlogic.Connect()
	},
}

func init() {
	rootCmd.AddCommand(connectCmd)

	// Here you will define your flags and configuration settings.
	connectCmd.Flags().StringVarP(&server, "server", "s", "", "Server URL to connect")
	connectCmd.Flags().Int64VarP(&port, "port", "P", 0, "port to the mqtt server")
	connectCmd.Flags().StringVarP(&user, "user", "", "", "the user account")
	connectCmd.Flags().StringVarP(&password, "password", "", "", "the password for the account")
	connectCmd.Flags().StringSliceVar(&topics, "topic", []string{}, "the topics to be subscribed on the mqtt server")
	connectCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "set verbose output (optional)")
	connectCmd.MarkFlagRequired("server")
	connectCmd.MarkFlagRequired("port")
	connectCmd.MarkFlagRequired("user")
	connectCmd.MarkFlagRequired("pass")
	connectCmd.MarkFlagRequired("topics")

}
