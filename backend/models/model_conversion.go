// File generated via python script because I'm no masochist

package models

import "msgr/database"

func InsertChatParamsToSqlc(insertChatParams InsertChatParams) database.InsertChatParams {
	return database.InsertChatParams{
		ID:         ToPgtypeUUID(insertChatParams.ID),
		FirstUser:  ToPgtypeUUID(insertChatParams.FirstUser),
		SecondUser: ToPgtypeUUID(insertChatParams.SecondUser),
	}
}

func InsertMessageParamsToSqlc(insertMessageParams InsertMessageParams) database.InsertMessageParams {
	return database.InsertMessageParams{
		ID:       ToPgtypeUUID(insertMessageParams.ID),
		Chat:     ToPgtypeUUID(insertMessageParams.Chat),
		Sender:   ToPgtypeUUID(insertMessageParams.Sender),
		Receiver: ToPgtypeUUID(insertMessageParams.Receiver),
		Message:  insertMessageParams.Message,
	}
}

func InsertUserParamsToSqlc(insertUserParams InsertUserParams) database.InsertUserParams {
	return database.InsertUserParams{
		ID:       ToPgtypeUUID(insertUserParams.ID),
		Name:     insertUserParams.Name,
		Password: insertUserParams.Password,
		Email:    insertUserParams.Email,
	}
}

func ChatToSqlc(chat Chat) database.Chat {
	return database.Chat{
		ID:         ToPgtypeUUID(chat.ID),
		FirstUser:  ToPgtypeUUID(chat.FirstUser),
		SecondUser: ToPgtypeUUID(chat.SecondUser),
		CreatedAt:  ToPgtypeTimestamp(chat.CreatedAt),
	}
}

func MessageToSqlc(message Message) database.Message {
	return database.Message{
		ID:       ToPgtypeUUID(message.ID),
		Chat:     ToPgtypeUUID(message.Chat),
		Sender:   ToPgtypeUUID(message.Sender),
		Receiver: ToPgtypeUUID(message.Receiver),
		Message:  message.Message,
		SentAt:   ToPgtypeTimestamp(message.SentAt),
	}
}

func UserToSqlc(user User) database.User {
	return database.User{
		ID:           ToPgtypeUUID(user.ID),
		Name:         user.Name,
		Password:     user.Password,
		Email:        user.Email,
		RegisteredAt: ToPgtypeTimestamp(user.RegisteredAt),
	}
}

func InsertChatParamsFromSqlc(insertChatParams database.InsertChatParams) InsertChatParams {
	return InsertChatParams{
		ID:         ToGoogleUUID(insertChatParams.ID),
		FirstUser:  ToGoogleUUID(insertChatParams.FirstUser),
		SecondUser: ToGoogleUUID(insertChatParams.SecondUser),
	}
}

func InsertMessageParamsFromSqlc(insertMessageParams database.InsertMessageParams) InsertMessageParams {
	return InsertMessageParams{
		ID:       ToGoogleUUID(insertMessageParams.ID),
		Chat:     ToGoogleUUID(insertMessageParams.Chat),
		Sender:   ToGoogleUUID(insertMessageParams.Sender),
		Receiver: ToGoogleUUID(insertMessageParams.Receiver),
		Message:  insertMessageParams.Message,
	}
}

func InsertUserParamsFromSqlc(insertUserParams database.InsertUserParams) InsertUserParams {
	return InsertUserParams{
		ID:       ToGoogleUUID(insertUserParams.ID),
		Name:     insertUserParams.Name,
		Password: insertUserParams.Password,
		Email:    insertUserParams.Email,
	}
}

func ChatFromSqlc(chat database.Chat) Chat {
	return Chat{
		ID:         ToGoogleUUID(chat.ID),
		FirstUser:  ToGoogleUUID(chat.FirstUser),
		SecondUser: ToGoogleUUID(chat.SecondUser),
		CreatedAt:  ToTime(chat.CreatedAt),
	}
}

func MessageFromSqlc(message database.Message) Message {
	return Message{
		ID:       ToGoogleUUID(message.ID),
		Chat:     ToGoogleUUID(message.Chat),
		Sender:   ToGoogleUUID(message.Sender),
		Receiver: ToGoogleUUID(message.Receiver),
		Message:  message.Message,
		SentAt:   ToTime(message.SentAt),
	}
}

func UserFromSqlc(user database.User) User {
	return User{
		ID:           ToGoogleUUID(user.ID),
		Name:         user.Name,
		Password:     user.Password,
		Email:        user.Email,
		RegisteredAt: ToTime(user.RegisteredAt),
	}
}
