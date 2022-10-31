package main

func onLoginRequest(client *ClientConnection, command LoginRequestCommand) {
	response := LoginDeniedCommand{
		reason: communicationProblem,
	}
	loginDenied(client, response)
}
