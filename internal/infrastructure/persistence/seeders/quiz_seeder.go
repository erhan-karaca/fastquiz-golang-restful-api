package seeders

import (
	"fastquiz-api/internal/domain/constants"
	"fastquiz-api/internal/domain/entities"
	"fastquiz-api/internal/infrastructure/persistence/db"
	"fastquiz-api/pkg/utils"
	"log"
)

func SeedQuizzes() {

	movies := []struct {
		Name     string
		SourceID string
	}{
		{"The Shawshank Redemption", "tt0111161"},
		{"The Godfather", "tt0068646"},
		{"The Dark Knight", "tt0468569"},
		{"The Godfather Part II", "tt0071562"},
		{"12 Angry Men", "tt0050083"},
		{"Schindler's List", "tt0108052"},
		{"The Lord of the Rings: The Return of the King", "tt0167260"},
		{"Pulp Fiction", "tt0110912"},
		{"The Lord of the Rings: The Fellowship of the Ring", "tt0120737"},
		{"The Good, the Bad and the Ugly", "tt0060196"},
		{"Forrest Gump", "tt0109830"},
		{"Fight Club", "tt0137523"},
		{"Inception", "tt1375666"},
		{"The Lord of the Rings: The Two Towers", "tt0167261"},
		{"Star Wars: Episode V - The Empire Strikes Back", "tt0080684"},
		{"The Matrix", "tt0133093"},
		{"Goodfellas", "tt0099685"},
		{"One Flew Over the Cuckoo's Nest", "tt0073486"},
		{"Seven Samurai", "tt0047478"},
		{"Se7en", "tt0114369"},
		{"It's a Wonderful Life", "tt0038650"},
		{"The Silence of the Lambs", "tt0102926"},
		{"City of God", "tt0317248"},
		{"Saving Private Ryan", "tt0120815"},
		{"Interstellar", "tt0816692"},
		{"Life Is Beautiful", "tt0118799"},
		{"The Green Mile", "tt0120689"},
		{"Star Wars: Episode IV - A New Hope", "tt0076759"},
		{"Terminator 2: Judgment Day", "tt0103064"},
		{"Back to the Future", "tt0088763"},
		{"Spirited Away", "tt0245429"},
		{"The Pianist", "tt0253474"},
		{"The Departed", "tt0407887"},
		{"Gladiator", "tt0172495"},
		{"The Prestige", "tt0482571"},
		{"Whiplash", "tt2582802"},
		{"Parasite", "tt6751668"},
		{"The Lion King", "tt0110357"},
		{"The Usual Suspects", "tt0114814"},
		{"The Avengers", "tt0848228"},
		{"The Lives of Others", "tt0405094"},
		{"Grave of the Fireflies", "tt0095327"},
		{"American History X", "tt0120586"},
		{"The Wolf of Wall Street", "tt0993846"},
		{"Casablanca", "tt0034583"},
		{"Once Upon a Time in the West", "tt0064116"},
		{"Django Unchained", "tt1853728"},
		{"The Shining", "tt0081505"},
		{"Avengers: Infinity War", "tt4154756"},
		{"Coco", "tt2380307"},
		{"Joker", "tt7286456"},
		{"WALL-E", "tt0910970"},
		{"Alien", "tt0078748"},
		{"The Dark Knight Rises", "tt1345836"},
		{"Memento", "tt0209144"},
		{"Apocalypse Now", "tt0078788"},
		{"Princess Mononoke", "tt0119698"},
		{"Dr. Strangelove", "tt0057012"},
		{"Oldboy", "tt0364569"},
		{"The Great Dictator", "tt0032553"},
		{"Cinema Paradiso", "tt0095765"},
		{"Braveheart", "tt0112573"},
		{"Aliens", "tt0090605"},
		{"The Hunt", "tt2106476"},
		{"Amélie", "tt0211915"},
		{"A Clockwork Orange", "tt0066921"},
		{"Scarface", "tt0086250"},
		{"Taxi Driver", "tt0075314"},
		{"The Grand Budapest Hotel", "tt2278388"},
		{"Toy Story", "tt0114709"},
		{"Inglourious Basterds", "tt0361748"},
		{"Up", "tt1049413"},
		{"The Social Network", "tt1285016"},
		{"The Incredibles", "tt0317705"},
		{"Monsters, Inc.", "tt0198781"},
		{"Shrek", "tt0126029"},
		{"The Lego Movie", "tt1490017"},
		{"Finding Nemo", "tt0266543"},
		{"Ratatouille", "tt0382932"},
		{"Wreck-It Ralph", "tt1772341"},
		{"Zootopia", "tt2948356"},
		{"Frozen", "tt2294629"},
		{"Moana", "tt3521164"},
		{"Beauty and the Beast", "tt2771200"},
		{"Aladdin", "tt0103639"},
		{"The Little Mermaid", "tt0097757"},
		{"The Lion King (2019)", "tt6105098"},
		{"Snow White and the Seven Dwarfs", "tt0029583"},
		{"The Jungle Book", "tt3040964"},
		{"101 Dalmatians", "tt0055254"},
		{"Sleeping Beauty", "tt0053285"},
		{"Cinderella", "tt0042332"},
	}

	// Seed data
	quizzes := []entities.Quiz{}

	for _, movie := range movies {
		quiz := entities.Quiz{
			Name:       movie.Name,
			TypeID:     1,
			Status:     true,
			Language:   constants.Languages.Turkish,
			SourceType: "imdb",
			SourceID:   movie.SourceID,
			Slug:       utils.FormatSlug(movie.Name, "quiz"),
			Action:     entities.QuizCreated,
		}
		quizzes = append(quizzes, quiz)
	}

	for _, e := range quizzes {
		err := db.DB.Where("slug = ?", e.Slug).FirstOrCreate(&e).Error
		if err != nil {
			log.Printf("Failed to seed quiz: %s, error: %v", e.Name, err)
		} else {
			log.Printf("Seeded quiz: %s", e.Name)
		}
	}
}
