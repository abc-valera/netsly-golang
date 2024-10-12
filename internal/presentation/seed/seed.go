package seed

import (
	"context"

	"github.com/abc-valera/netsly-golang/internal/domain/entity"
	"github.com/abc-valera/netsly-golang/internal/domain/global"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query/selector"
	"github.com/abc-valera/netsly-golang/internal/domain/util/coderr"
)

// Seed is used to populate the database with initial data.
// Stops the program execution if an error occurs.
func Seed(entities entity.Entities) {
	ctx := context.Background()

	params := selector.Selector{
		Limit: 100,
	}

	userCreateReqs := []entity.UserCreateRequest{
		{
			Username: "abc_Valera",
			Email:    "abc_valera@example.com",
			Password: "abc_valera_password",
			Fullname: "Valeriy James",
			Status:   "Happy",
		},
		{
			Username: "StarGazer2024",
			Email:    "sj_stargazer@cosmicmail.com",
			Password: "Uy5*bN8@kR2!",
			Fullname: "Sarah Johnson",
			Status:   "Excited",
		},
		{
			Username: "TechNinja42",
			Email:    "mikeB_ninja@techworld.io",
			Password: "Jh6^fD4%wS9#",
			Fullname: "Michael Brown",
			Status:   "Busy",
		},
		{
			Username: "BookwormEmma",
			Email:    "emma.reads@literarynet.org",
			Password: "Lm3$tG7!xW1@",
			Fullname: "Emma Wilson",
			Status:   "Relaxed",
		},
		{
			Username: "MountainExplorer",
			Email:    "alex.peaks@adventuremail.com",
			Password: "Qs9#hY2@fK5!",
			Fullname: "Alexander Lee",
			Status:   "Focused",
		},
	}
	for _, user := range userCreateReqs {
		coderr.Must(entities.User.Create(ctx, user))
	}
	users := coderr.Must(entities.User.GetAll(ctx, params))

	jokeCreateReqs := []entity.JokeCreateRequest{
		{
			Title:       "The Chicken and the Egg",
			Text:        "Which came first, the chicken or the egg? The rooster.",
			Explanation: "The joke is that the rooster is the one who lays the eggs, not the hen.",
			UserID:      users[0].ID,
		},
		{
			Title:       "Math Humor",
			Text:        "Why was six afraid of seven? Because seven eight nine!",
			Explanation: "The joke plays on the similarity between 'ate' and 'eight' in pronunciation.",
			UserID:      users[0].ID,
		},
		{
			Title:       "Vegetable Pun",
			Text:        "What did the carrot say to the wheat? Lettuce rest, I'm feeling beat!",
			Explanation: "This joke uses puns with 'lettuce' sounding like 'let us' and 'beat' referring to both tiredness and a vegetable.",
			UserID:      users[1].ID,
		},
		{
			Title:       "Computer Bug",
			Text:        "Why do programmers prefer dark mode? Because light attracts bugs!",
			Explanation: "This joke plays on the dual meaning of 'bugs' as both insects and programming errors.",
			UserID:      users[1].ID,
		},
		{
			Title:       "Musical Fruit",
			Text:        "What kind of music do planets listen to? Neptunes!",
			Explanation: "The joke is a pun on 'Neptune' (a planet) and 'tunes' (music).",
			UserID:      users[2].ID,
		},
		{
			Title:       "Cheese Joke",
			Text:        "What do you call cheese that isn't yours? Nacho cheese!",
			Explanation: "This is a play on words, where 'nacho cheese' sounds like 'not your cheese'.",
			UserID:      users[3].ID,
		},
		{
			Title:       "Time Flies",
			Text:        "Why did the man throw the clock out the window? He wanted to see time fly!",
			Explanation: "This joke plays on the idiom 'time flies' by literally making a clock fly.",
			UserID:      users[3].ID,
		},
		{
			Title:       "Invisible Man",
			Text:        "I used to be addicted to soap, but I'm clean now.",
			Explanation: "This joke uses the dual meaning of 'clean' as both free from addiction and physically clean.",
			UserID:      users[3].ID,
		},
		{
			Title:       "Pirate Math",
			Text:        "What's a pirate's favorite letter? You might think it's R, but his true love be the C.",
			Explanation: "This joke plays on the pirate accent ('arr') and the sea ('C') that pirates sail on.",
			UserID:      users[4].ID,
		},
		{
			Title:       "Tree Humor",
			Text:        "Why did the tree go to the dentist? To get a root canal!",
			Explanation: "This joke connects the dental procedure 'root canal' with the roots of a tree.",
			UserID:      users[4].ID,
		},
	}
	for _, joke := range jokeCreateReqs {
		coderr.Must(entities.Joke.Create(ctx, joke))
	}
	jokes := coderr.Must(entities.Joke.SearchAllByTitle(ctx, "", params))

	likeCreateReqs := []entity.LikeCreateRequest{
		{
			UserID: users[1].ID,
			JokeID: jokes[0].ID,
		},
		{
			UserID: users[0].ID,
			JokeID: jokes[2].ID,
		},
		{
			UserID: users[3].ID,
			JokeID: jokes[5].ID,
		},
		{
			UserID: users[2].ID,
			JokeID: jokes[1].ID,
		},
		{
			UserID: users[4].ID,
			JokeID: jokes[7].ID,
		},
		{
			UserID: users[1].ID,
			JokeID: jokes[3].ID,
		},
		{
			UserID: users[0].ID,
			JokeID: jokes[8].ID,
		},
		{
			UserID: users[2].ID,
			JokeID: jokes[4].ID,
		},
		{
			UserID: users[3].ID,
			JokeID: jokes[9].ID,
		},
		{
			UserID: users[4].ID,
			JokeID: jokes[6].ID,
		},
		{
			UserID: users[1].ID,
			JokeID: jokes[5].ID,
		},
		{
			UserID: users[0].ID,
			JokeID: jokes[7].ID,
		},
		{
			UserID: users[2].ID,
			JokeID: jokes[9].ID,
		},
		{
			UserID: users[3].ID,
			JokeID: jokes[1].ID,
		},
		{
			UserID: users[4].ID,
			JokeID: jokes[3].ID,
		},
		{
			UserID: users[1].ID,
			JokeID: jokes[8].ID,
		},
		{
			UserID: users[0].ID,
			JokeID: jokes[4].ID,
		},
		{
			UserID: users[2].ID,
			JokeID: jokes[6].ID,
		},
		{
			UserID: users[3].ID,
			JokeID: jokes[2].ID,
		},
		{
			UserID: users[4].ID,
			JokeID: jokes[0].ID,
		},
	}
	for _, like := range likeCreateReqs {
		coderr.Must(entities.Like.Create(ctx, like))
	}

	commentCreateReqs := []entity.CommentCreateRequest{
		{
			Text:   "Haha, that's a good one!",
			UserID: users[1].ID,
			JokeID: jokes[0].ID,
		},
		{
			Text:   "I can't stop laughing!",
			UserID: users[3].ID,
			JokeID: jokes[2].ID,
		},
		{
			Text:   "Clever wordplay, love it.",
			UserID: users[0].ID,
			JokeID: jokes[5].ID,
		},
		{
			Text:   "This one's my favorite so far.",
			UserID: users[2].ID,
			JokeID: jokes[8].ID,
		},
		{
			Text:   "I'm going to use this at work!",
			UserID: users[4].ID,
			JokeID: jokes[1].ID,
		},
		{
			Text:   "Didn't get it at first, but now I'm chuckling.",
			UserID: users[1].ID,
			JokeID: jokes[7].ID,
		},
		{
			Text:   "Classic dad joke material right here.",
			UserID: users[3].ID,
			JokeID: jokes[4].ID,
		},
		{
			Text:   "My kids are going to love this one!",
			UserID: users[0].ID,
			JokeID: jokes[9].ID,
		},
		{
			Text:   "Groan-worthy, but in a good way.",
			UserID: users[2].ID,
			JokeID: jokes[3].ID,
		},
		{
			Text:   "I snorted my coffee reading this. Thanks!",
			UserID: users[4].ID,
			JokeID: jokes[6].ID,
		},
		{
			Text:   "This joke is out of this world!",
			UserID: users[1].ID,
			JokeID: jokes[2].ID,
		},
		{
			Text:   "I'm stealing this for my next party.",
			UserID: users[3].ID,
			JokeID: jokes[0].ID,
		},
		{
			Text:   "You had me in the first half, not gonna lie.",
			UserID: users[0].ID,
			JokeID: jokes[7].ID,
		},
		{
			Text:   "This joke is so cheesy, I love it!",
			UserID: users[2].ID,
			JokeID: jokes[5].ID,
		},
		{
			Text:   "I can't decide if this is brilliant or terrible. Maybe both?",
			UserID: users[4].ID,
			JokeID: jokes[8].ID,
		},
		{
			Text:   "This joke just made my day better.",
			UserID: users[1].ID,
			JokeID: jokes[4].ID,
		},
		{
			Text:   "I'm going to annoy my family with this one all week.",
			UserID: users[3].ID,
			JokeID: jokes[9].ID,
		},
		{
			Text:   "Short, sweet, and hilarious!",
			UserID: users[0].ID,
			JokeID: jokes[1].ID,
		},
		{
			Text:   "I didn't expect that punchline at all. Well done!",
			UserID: users[2].ID,
			JokeID: jokes[6].ID,
		},
		{
			Text:   "This joke is so bad it's good.",
			UserID: users[4].ID,
			JokeID: jokes[3].ID,
		},
	}
	for _, comment := range commentCreateReqs {
		coderr.Must(entities.Comment.Create(ctx, comment))
	}

	roomCreateReqs := []entity.RoomCreateRequest{
		{
			Name:          "Joke Room",
			Description:   "A room for sharing jokes and having a good laugh.",
			CreatorUserID: users[0].ID,
		},
		{
			Name:          "Pun Paradise",
			Description:   "Where wordplay reigns supreme. Enter at your own risk!",
			CreatorUserID: users[1].ID,
		},
		{
			Name:          "Stand-Up Corner",
			Description:   "Share your best stand-up comedy material and get feedback.",
			CreatorUserID: users[2].ID,
		},
		{
			Name:          "Meme Madness",
			Description:   "For those who prefer their humor in image format. Memes welcome!",
			CreatorUserID: users[3].ID,
		},
		{
			Name:          "Dad Jokes Central",
			Description:   "The perfect place for all those groan-worthy dad jokes.",
			CreatorUserID: users[4].ID,
		},
	}
	for _, room := range roomCreateReqs {
		coderr.Must(entities.Room.Create(ctx, room))
	}
	rooms := coderr.Must(entities.Room.SearchAllByName(ctx, "", params))

	roomMembersReqs := []entity.RoomMemberCreateRequest{
		{
			UserID: users[1].ID,
			RoomID: rooms[0].ID,
		},
		{
			UserID: users[0].ID,
			RoomID: rooms[1].ID,
		},
		{
			UserID: users[2].ID,
			RoomID: rooms[0].ID,
		},
		{
			UserID: users[3].ID,
			RoomID: rooms[1].ID,
		},
		{
			UserID: users[4].ID,
			RoomID: rooms[2].ID,
		},
		{
			UserID: users[0].ID,
			RoomID: rooms[3].ID,
		},
		{
			UserID: users[1].ID,
			RoomID: rooms[4].ID,
		},
		{
			UserID: users[3].ID,
			RoomID: rooms[0].ID,
		},
		{
			UserID: users[4].ID,
			RoomID: rooms[1].ID,
		},
		{
			UserID: users[0].ID,
			RoomID: rooms[2].ID,
		},
		{
			UserID: users[1].ID,
			RoomID: rooms[3].ID,
		},
		{
			UserID: users[2].ID,
			RoomID: rooms[4].ID,
		},
		{
			UserID: users[4].ID,
			RoomID: rooms[0].ID,
		},
		{
			UserID: users[0].ID,
			RoomID: rooms[4].ID,
		},
		{
			UserID: users[1].ID,
			RoomID: rooms[2].ID,
		},
		{
			UserID: users[2].ID,
			RoomID: rooms[3].ID,
		},
		{
			UserID: users[3].ID,
			RoomID: rooms[4].ID,
		},
	}
	for _, roomMember := range roomMembersReqs {
		coderr.Must(entities.RoomMember.Create(ctx, roomMember))
	}

	roomMessages := []entity.RoomMessageCreateRequest{
		{
			Text:   "Welcome to the Joke Room! Let's get this party started!",
			RoomID: rooms[0].ID,
			UserID: users[0].ID,
		},
		{
			Text:   "Hey everyone! Who's ready for some wordplay?",
			RoomID: rooms[1].ID,
			UserID: users[2].ID,
		},
		{
			Text:   "Stand-up Corner is now open! Share your best material!",
			RoomID: rooms[2].ID,
			UserID: users[1].ID,
		},
		{
			Text:   "Meme lovers unite! Post your freshest memes here.",
			RoomID: rooms[3].ID,
			UserID: users[3].ID,
		},
		{
			Text:   "Dad joke incoming: Why don't scientists trust atoms? Because they make up everything!",
			RoomID: rooms[4].ID,
			UserID: users[4].ID,
		},
		{
			Text:   "I've got a great joke about construction, but I'm still working on it.",
			RoomID: rooms[0].ID,
			UserID: users[1].ID,
		},
		{
			Text:   "Time flies like an arrow. Fruit flies like a banana.",
			RoomID: rooms[1].ID,
			UserID: users[0].ID,
		},
		{
			Text:   "Why did the scarecrow win an award? He was outstanding in his field!",
			RoomID: rooms[4].ID,
			UserID: users[2].ID,
		},
		{
			Text:   "Just uploaded a new meme! Check it out and let me know what you think.",
			RoomID: rooms[3].ID,
			UserID: users[4].ID,
		},
		{
			Text:   "What do you call a fake noodle? An impasta!",
			RoomID: rooms[0].ID,
			UserID: users[3].ID,
		},
		{
			Text:   "I'm working on a new bit about procrastination, but I'll tell you later.",
			RoomID: rooms[2].ID,
			UserID: users[0].ID,
		},
		{
			Text:   "Why did the math book look so sad? Because it had too many problems.",
			RoomID: rooms[1].ID,
			UserID: users[4].ID,
		},
		{
			Text:   "Just heard a great joke about amnesia, but I forgot how it goes.",
			RoomID: rooms[0].ID,
			UserID: users[2].ID,
		},
		{
			Text:   "Why don't skeletons fight each other? They don't have the guts.",
			RoomID: rooms[4].ID,
			UserID: users[1].ID,
		},
		{
			Text:   "New meme format alert! Get ready for some fresh content.",
			RoomID: rooms[3].ID,
			UserID: users[0].ID,
		},
	}
	for _, roomMessage := range roomMessages {
		coderr.Must(entities.RoomMessage.Create(ctx, roomMessage))
	}

	global.Log().Info("Database has been seeded")
}
