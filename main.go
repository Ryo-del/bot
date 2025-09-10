package main

import (
	"log"

	tgbotapi "github.com/ilpy20/telegram-bot-api/v7"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("7582375690:AAHOzDEALtSjbiovqytc6cj0p1xwj_m_0nU")
	if err != nil {
		log.Panic(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		// Обработка команды /start
		if update.Message != nil && update.Message.Text == "/start" {
			// Создаем WebApp кнопку
			webApp := tgbotapi.WebAppInfo{
				URL: "https://ryo-del.github.io/dzshka/",
			}
			buttonWebApp := tgbotapi.NewInlineKeyboardButtonWebApp("Открыть", webApp)

			// Кнопка FAQ
			buttonFaq := tgbotapi.NewInlineKeyboardButtonData("FAQ", "/faq")

			// Inline клавиатура
			keyboard := tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(buttonWebApp),
				tgbotapi.NewInlineKeyboardRow(buttonFaq),
			)

			// Отправляем сообщение с inline-кнопками
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет! Я — Dz Bot. Здесь ты можешь узнать своё домашнее задание.")
			msg.ReplyMarkup = keyboard
			bot.Send(msg)
		}

		// Обработка нажатия кнопки FAQ
		if update.CallbackQuery != nil && update.CallbackQuery.Data == "/faq" {
			// Удаляем предыдущее сообщение
			bot.Request(tgbotapi.NewDeleteMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID))

			// Новое сообщение с кнопкой "Как добавлять дз?"
			buttonDz := tgbotapi.NewInlineKeyboardButtonData("Как добавлять дз?", "/add_dz")
			buttonEr := tgbotapi.NewInlineKeyboardButtonData("Нашёл баг?", "/errors")
			buttonnotwork := tgbotapi.NewInlineKeyboardButtonData("бот перестал работать?", "/notwork")
			buttontheme := tgbotapi.NewInlineKeyboardButtonData("Как поменять тему?", "/theme")
			buttondonate := tgbotapi.NewInlineKeyboardButtonData("Хочу поддержать проект!", "/donate")
			buttonweb := tgbotapi.NewInlineKeyboardButtonData("WebApp не работает?", "/web")
			buttonnodz := tgbotapi.NewInlineKeyboardButtonData("Дз не отображается после добавления?", "/nodz")
			buttonexit := tgbotapi.NewInlineKeyboardButtonData("Назад", "/exit")
			keyboard := tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(buttonDz),
				tgbotapi.NewInlineKeyboardRow(buttonEr),
				tgbotapi.NewInlineKeyboardRow(buttonnotwork),
				tgbotapi.NewInlineKeyboardRow(buttontheme),
				tgbotapi.NewInlineKeyboardRow(buttonweb),
				tgbotapi.NewInlineKeyboardRow(buttonnodz),
				tgbotapi.NewInlineKeyboardRow(buttondonate),
				tgbotapi.NewInlineKeyboardRow(buttonexit),
			)
			newMsg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Выберите вопрос:")
			newMsg.ReplyMarkup = keyboard
			bot.Send(newMsg)

			// Убираем "часики" на кнопке
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "")
			bot.Request(callback)
		}
		if update.CallbackQuery != nil && update.CallbackQuery.Data == "/add_dz" {
			bot.Request(tgbotapi.NewDeleteMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID))
			newMsg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Чтобы добавить домашнее задание, откройте WebApp, выберите предмет, нажмите на кнопку 'редактировать' и введите задание. После этого нажмите 'Сохранить'.")
			bot.Send(newMsg)
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "")
			bot.Request(callback)

		}
		if update.CallbackQuery != nil && update.CallbackQuery.Data == "/errors" {
			bot.Request(tgbotapi.NewDeleteMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID))
			newMsg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Если вы нашли баг, то опишите его и пришлите скрин мне в лс(@Hirasawaaaa)")
			bot.Send(newMsg)
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "")
			bot.Request(callback)
		}
		if update.CallbackQuery != nil && update.CallbackQuery.Data == "/notwork" {
			bot.Request(tgbotapi.NewDeleteMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID))
			newMsg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Бот может не отвечать на сообщения по нескольким причинам:\n\n1.Проблемы с Хостингом - ")
			bot.Send(newMsg)
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "")
			bot.Request(callback)
		}
	}
}
