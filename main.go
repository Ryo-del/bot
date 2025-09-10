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
		chatID := int64(0)
		if update.Message != nil {
			chatID = update.Message.Chat.ID
		} else if update.CallbackQuery != nil {
			chatID = update.CallbackQuery.Message.Chat.ID
		}

		if update.Message != nil && update.Message.Text == "/new" {
			bot.Send(tgbotapi.NewMessage(chatID, "update 1.0.0\n\n В планах:\n"+
				"- Добавление возможности выбирать тему бота и загружать сторонние темы.\n"+
				"- Поддержка работы с более крупными файлами — до 100 МБ.\n"+
				"- Система добавления домашнего задания с привязкой к конкретным датам.\n"+
				"- Готовое решение от ChatGPT.\n"+
				"- И другие улучшения для удобства использования бота."))

		}

		// --- Обработка команды /start ---
		if (update.Message != nil && update.Message.Text == "/start") ||
			(update.CallbackQuery != nil && update.CallbackQuery.Data == "/start") {

			// Создаем WebApp кнопку
			webApp := tgbotapi.WebAppInfo{
				URL: "https://ryo-del.github.io/dzshka/",
			}
			buttonWebApp := tgbotapi.NewInlineKeyboardButtonWebApp("Открыть", webApp)
			buttonFaq := tgbotapi.NewInlineKeyboardButtonData("FAQ", "/faq")
			keyboard := tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(buttonWebApp),
				tgbotapi.NewInlineKeyboardRow(buttonFaq),
			)

			msg := tgbotapi.NewMessage(chatID, "Привет! Я — Dz Bot. Здесь ты можешь узнать своё домашнее задание.")
			msg.ReplyMarkup = keyboard
			bot.Send(msg)

			if update.CallbackQuery != nil {
				callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "")
				bot.Request(callback)
			}
		}

		// --- Обработка коллбэков ---
		if update.CallbackQuery != nil {
			chatID := update.CallbackQuery.Message.Chat.ID

			switch update.CallbackQuery.Data {
			case "/faq":
				bot.Request(tgbotapi.NewDeleteMessage(chatID, update.CallbackQuery.Message.MessageID))
				buttonDz := tgbotapi.NewInlineKeyboardButtonData("Как добавлять дз?", "/add_dz")
				buttonEr := tgbotapi.NewInlineKeyboardButtonData("Нашёл баг?", "/errors")
				buttonnotwork := tgbotapi.NewInlineKeyboardButtonData("Бот перестал работать?", "/notwork")
				buttontheme := tgbotapi.NewInlineKeyboardButtonData("Как поменять тему?", "/theme")
				buttondonate := tgbotapi.NewInlineKeyboardButtonData("Хочу поддержать проект!", "/donate")
				buttonweb := tgbotapi.NewInlineKeyboardButtonData("WebApp не работает?", "/web")
				buttonnodz := tgbotapi.NewInlineKeyboardButtonData("Дз не отображается после добавления?", "/nodz")
				buttonexit := tgbotapi.NewInlineKeyboardButtonData("Назад", "/start")
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
				newMsg := tgbotapi.NewMessage(chatID, "Выберите вопрос:")
				newMsg.ReplyMarkup = keyboard
				bot.Send(newMsg)
				bot.Request(tgbotapi.NewCallback(update.CallbackQuery.ID, ""))
			case "/add_dz":
				bot.Send(tgbotapi.NewMessage(chatID, "Чтобы добавить домашнее задание, откройте WebApp, выберите предмет, нажмите на кнопку 'редактировать' и введите задание. После этого нажмите 'Сохранить'."))
			case "/errors":
				bot.Send(tgbotapi.NewMessage(chatID, "Если вы нашли баг, то опишите его и пришлите скрин мне в лс(@Hirasawaaaa)"))
			case "/notwork":
				bot.Send(tgbotapi.NewMessage(chatID, "Бот может не отвечать на сообщения по нескольким причинам:\n\n1.Проблемы с Хостингом - иногда серверы, на которых размещены боты, могут испытывать технические трудности или быть временно недоступными.\n\n2.Технические работы - разработка бота может проводить обновления или технические работы, что может привести к временному отключению бота.\n\nПроверяйте описание бота, там я пишу, если бот не работает по какой-то причине."))
			case "/theme":
				bot.Send(tgbotapi.NewMessage(chatID, "Вы можете полностью поменять стиль сайта WebApp. Инструкция и исходники тут: https://github.com/Ryo-del/dzshka"))
			case "/donate":
				bot.Send(tgbotapi.NewMessage(chatID, "Если вы хотите поддержать проект, вы можете отправить донат на Tinkoff: 2200701989137176. Спасибо за вашу поддержку!"))
			case "/web":
				bot.Send(tgbotapi.NewMessage(chatID, "Если WebApp не открывается, на это может быть несколько причин:\n\n1.Проверьте не включен ли у вас vpn.\n2.Проблемы с хостингом\n3.Технические работы\n\nЕсли в описании бота нет информации о том, что WebApp не работает, то напишите мне в лс(@Hirasawaaaa), я постараюсь помочь."))
			case "/nodz":
				bot.Send(tgbotapi.NewMessage(chatID, "Если дз не отображается после добавления, попробуйте перезапустить WebApp или подождать от 5 до 30 секунд. Так же проверьте, не превышает ли файлы размер 10мб-20мб. Если проблема не решится, напишите мне в лс(@Hirasawaaaa), я постараюсь помочь."))
			}
		}
	}
}
