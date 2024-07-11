package seed

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/internal/core/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/domain"
	"github.com/abc-valera/netsly-api-golang/internal/domain/entity"
	"github.com/abc-valera/netsly-api-golang/internal/domain/entityTransactor"
	"github.com/abc-valera/netsly-api-golang/internal/domain/global"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query/selector"
)

// Seed is used to populate the database with initial data.
// Stops the program execution if an error occurs.
func Seed(entities domain.Entities, entityTransactor entityTransactor.ITransactor) {
	params := selector.Selector{
		Order: "asc",
		Limit: 100,
	}

	txFunc := func(ctx context.Context, txEntities domain.Entities) error {
		// Users
		userRequests := []entity.UserCreateRequest{
			{
				Username: "Valera",
				Email:    "valera@gmail.com",
				Password: "valera",
				Fullname: "Valeriy",
				Status:   "Working :(",
			},
			{
				Username: "John",
				Email:    "john@gmail.com",
				Password: "john",
				Fullname: "John Doe",
				Status:   "On vacation :)",
			},
			{
				Username: "Jane",
				Email:    "jane@gmail.com",
				Password: "jane",
				Fullname: "Jane Doe",
				Status:   "Working from home :)",
			},
			{
				Username: "Alice",
				Email:    "alice@gmail.com",
				Password: "alice",
				Fullname: "Alice Wonderland",
				Status:   "On sick leave :(",
			},
			{
				Username: "Bob",
				Email:    "bob@gmail.com",
				Password: "bob",
				Fullname: "Bob Builder",
				Status:   "In a meeting :|",
			},
			{
				Username: "Charlie",
				Email:    "charlie@gmail.com",
				Password: "charlie",
				Fullname: "Charlie Chaplin",
				Status:   "On a business trip :)",
			},
			{
				Username: "Dave",
				Email:    "dave@gmail.com",
				Password: "dave",
				Fullname: "Dave Grohl",
				Status:   "In the studio :)",
			},
			{
				Username: "Eve",
				Email:    "eve@gmail.com",
				Password: "eve",
				Fullname: "Eve Polastri",
				Status:   "On a mission :)",
			},
			{
				Username: "Frank",
				Email:    "frank@gmail.com",
				Password: "frank",
				Fullname: "Frank Sinatra",
				Status:   "Singing :)",
			},
			{
				Username: "Grace",
				Email:    "grace@gmail.com",
				Password: "grace",
				Fullname: "Grace Hopper",
				Status:   "Coding :)",
			},
		}
		for _, user := range userRequests {
			coderr.Must(entities.User.Create(context.Background(), user))
		}
		// Generated users
		users := coderr.Must(entities.User.SearchAllByUsername(context.Background(), "", params))

		// Jokes
		jokeRequests := []entity.JokeCreateRequest{
			{
				UserID:      users[0].ID,
				Title:       "Dad joke about a scarecrow",
				Text:        "Why did the scarecrow win an award? Because he was outstanding in his field!",
				Explanation: "The joke is a play on words. The scarecrow is outstanding in his field because he is in a field of crops, but also outstanding because he is excellent at his job of scaring crows away.",
			},
			{
				UserID:      users[0].ID,
				Title:       "Dad joke about a pizza",
				Text:        "Why did the tomato turn red? Because it saw the salad dressing!",
				Explanation: "The joke is a play on words. The tomato 'turns red' because it's blushing, as if it were embarrassed.",
			},
			{
				UserID:      users[1].ID,
				Title:       "Dad joke about a factory",
				Text:        "What do you call a factory that makes good products? A satisfactory!",
				Explanation: "The joke is a play on words. The word 'satisfactory' is a combination of 'satisfy' and 'factory'.",
			},
			{
				UserID:      users[1].ID,
				Title:       "Dad joke about a computer",
				Text:        "Why don't computers take their hats off? Because they have bad data caps!",
				Explanation: "The joke is a play on words. 'Data cap' is a term in computing, but in this joke it's taken literally as a type of hat.",
			},
			{
				UserID:      users[2].ID,
				Title:       "Dad joke about a book",
				Text:        "I told my wife she should embrace her mistakes. She gave me a hug.",
				Explanation: "The joke is a play on words. The word 'embrace' means to hug, but also means to accept something willingly.",
			},
			{
				UserID:      users[2].ID,
				Title:       "Dad joke about a chicken",
				Text:        "Why did the chicken go to the seance? To talk to the other side!",
				Explanation: "The joke is a play on the classic 'why did the chicken cross the road' joke. The 'other side' refers to the afterlife in the context of a seance, but also to the other side of the road.",
			},
			{
				UserID:      users[3].ID,
				Title:       "Dad joke about a ghost",
				Text:        "Why don't ghosts like rain? It dampens their spirits!",
				Explanation: "The joke is a play on words. 'Dampen one's spirits' is an idiom meaning to make someone less enthusiastic or happy, but in this joke it's taken literally as if the rain were affecting the ghosts' spiritual form.",
			},
			{
				UserID:      users[3].ID,
				Title:       "Dad joke about a sandwich",
				Text:        "I asked the librarian if the library had any books on paranoia. She whispered, 'They're right behind you.'",
				Explanation: "The joke is a play on words. The librarian is making a joke about paranoia, which is the fear that someone or something is out to get you.",
			},
			{
				UserID:      users[4].ID,
				Title:       "Dad joke about a tree",
				Text:        "Why do trees seem suspicious on sunny days? They just seem a bit shady!",
				Explanation: "The joke is a play on words. 'Shady' can mean suspicious, but in this context it's taken literally as providing shade from the sun.",
			},
			{
				UserID:      users[4].ID,
				Title:       "Dad joke about a calendar",
				Text:        "I told my wife she should embrace her mistakes. She gave me a hug.",
				Explanation: "The joke is a play on words. The word 'embrace' means to hug, but also means to accept something willingly.",
			},
			{
				UserID:      users[5].ID,
				Title:       "Dad joke about a pencil",
				Text:        "Why don't scientists trust atoms? Because they make up everything!",
				Explanation: "The joke is a play on words. Atoms are the basic units of matter and make up everything in the universe.",
			},
			{
				UserID:      users[6].ID,
				Title:       "Dad joke about a bicycle",
				Text:        "Why can't a bicycle stand up by itself? Because it's two-tired!",
				Explanation: "The joke is a play on words. 'Two-tired' sounds like 'too tired'.",
			},
			{
				UserID:      users[7].ID,
				Title:       "Dad joke about a clock",
				Text:        "Why did the man put his money in the blender? Because he wanted to make some liquid assets!",
				Explanation: "The joke is a play on words. 'Liquid assets' is a financial term, but in this joke it's taken literally.",
			},
			{
				UserID:      users[8].ID,
				Title:       "Dad joke about a cat",
				Text:        "Why don't cats play poker in the jungle? Too many cheetahs!",
				Explanation: "The joke is a play on words. 'Cheetahs' sounds like 'cheaters'.",
			},
			{
				UserID:      users[9].ID,
				Title:       "Dad joke about a fish",
				Text:        "Why don't fish play basketball? Because they're afraid of the net!",
				Explanation: "The joke is a play on words. 'Net' is a term used in both fishing and basketball.",
			},
		}
		for _, joke := range jokeRequests {
			coderr.Must(entities.Joke.Create(context.Background(), joke))
		}
		// Generated jokes
		jokes := coderr.Must(entities.Joke.SearchAllByTitle(context.Background(), "", params))

		// Likes
		likes := []entity.LikeCreateRequest{
			{
				UserID: users[0].ID,
				JokeID: jokes[0].ID,
			},
			{
				UserID: users[1].ID,
				JokeID: jokes[9].ID,
			},
			{
				UserID: users[2].ID,
				JokeID: jokes[4].ID,
			},
			{
				UserID: users[3].ID,
				JokeID: jokes[1].ID,
			},
			{
				UserID: users[4].ID,
				JokeID: jokes[7].ID,
			},
			{
				UserID: users[5].ID,
				JokeID: jokes[4].ID,
			},
			{
				UserID: users[6].ID,
				JokeID: jokes[2].ID,
			},
			{
				UserID: users[7].ID,
				JokeID: jokes[1].ID,
			},
			{
				UserID: users[8].ID,
				JokeID: jokes[6].ID,
			},
			{
				UserID: users[9].ID,
				JokeID: jokes[7].ID,
			},
			{
				UserID: users[0].ID,
				JokeID: jokes[1].ID,
			},
			{
				UserID: users[1].ID,
				JokeID: jokes[2].ID,
			},
			{
				UserID: users[2].ID,
				JokeID: jokes[3].ID,
			},
			{
				UserID: users[3].ID,
				JokeID: jokes[4].ID,
			},
			{
				UserID: users[4].ID,
				JokeID: jokes[5].ID,
			},
			{
				UserID: users[5].ID,
				JokeID: jokes[6].ID,
			},
			{
				UserID: users[6].ID,
				JokeID: jokes[7].ID,
			},
			{
				UserID: users[7].ID,
				JokeID: jokes[8].ID,
			},
			{
				UserID: users[8].ID,
				JokeID: jokes[9].ID,
			},
			{
				UserID: users[9].ID,
				JokeID: jokes[0].ID,
			},
		}
		for _, like := range likes {
			coderr.Must(entities.Like.Create(context.Background(), like))
		}

		// Comments
		comments := []entity.CommentCreateRequest{
			{
				UserID: users[0].ID,
				JokeID: jokes[2].ID,
				Text:   "Haha, that's a good one!",
			},
			{
				UserID: users[6].ID,
				JokeID: jokes[1].ID,
				Text:   "I don't get it.",
			},
			{
				UserID: users[1].ID,
				JokeID: jokes[0].ID,
				Text:   "That's hilarious!",
			},
			{
				UserID: users[2].ID,
				JokeID: jokes[1].ID,
				Text:   "I can't stop laughing!",
			},
			{
				UserID: users[3].ID,
				JokeID: jokes[2].ID,
				Text:   "This one made my day!",
			},
			{
				UserID: users[4].ID,
				JokeID: jokes[3].ID,
				Text:   "I'm going to tell this one at work.",
			},
			{
				UserID: users[5].ID,
				JokeID: jokes[4].ID,
				Text:   "I didn't see that punchline coming.",
			},
			{
				UserID: users[6].ID,
				JokeID: jokes[5].ID,
				Text:   "I don't get it. Can someone explain?",
			},
			{
				UserID: users[7].ID,
				JokeID: jokes[6].ID,
				Text:   "That's a classic!",
			},
			{
				UserID: users[8].ID,
				JokeID: jokes[7].ID,
				Text:   "I've heard this one before, but it's still funny.",
			},
			{
				UserID: users[9].ID,
				JokeID: jokes[8].ID,
				Text:   "This joke is so bad, it's good.",
			},
			{
				UserID: users[0].ID,
				JokeID: jokes[9].ID,
				Text:   "I laughed so hard I cried.",
			},
			{
				UserID: users[1].ID,
				JokeID: jokes[2].ID,
				Text:   "This is my new favorite joke.",
			},
			{
				UserID: users[2].ID,
				JokeID: jokes[3].ID,
				Text:   "I'm going to tell this one to my kids.",
			},
			{
				UserID: users[3].ID,
				JokeID: jokes[4].ID,
				Text:   "This one is a bit corny, but I like it.",
			},
			{
				UserID: users[4].ID,
				JokeID: jokes[5].ID,
				Text:   "I didn't get it at first, but now I do!",
			},
			{
				UserID: users[5].ID,
				JokeID: jokes[6].ID,
				Text:   "This one is really clever!",
			},
			{
				UserID: users[6].ID,
				JokeID: jokes[7].ID,
				Text:   "I love a good pun.",
			},
			{
				UserID: users[7].ID,
				JokeID: jokes[8].ID,
				Text:   "This one took me a second, but I got it.",
			},
			{
				UserID: users[8].ID,
				JokeID: jokes[9].ID,
				Text:   "I'm not usually a fan of dad jokes, but this one is good.",
			},
			{
				UserID: users[9].ID,
				JokeID: jokes[0].ID,
				Text:   "I can't wait to tell this one at the next family gathering.",
			},
			{
				UserID: users[0].ID,
				JokeID: jokes[1].ID,
				Text:   "This one is a bit over my head.",
			},
		}
		for _, comment := range comments {
			coderr.Must(entities.Comment.Create(context.Background(), comment))
		}

		// Rooms
		// roomRequests := []entity.RoomCreateRequest{
		// 	{
		// 		Name: "Dad Jokes",

		// 	},
		// 	{
		// 		Name: "Programming Puns",
		// 	},
		// 	{
		// 		Name: "Classic Jokes",
		// 	},
		// }

		return nil
	}

	coderr.NoErr(entityTransactor.PerformTX(context.Background(), txFunc))

	global.Log().Info("Database has been seeded")
}
