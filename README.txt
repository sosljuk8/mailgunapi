mailgun
    password = owlssowl81
    url = https://app.mailgun.com/app/account/setup
    username = andrew
    Key ID: 69a6bd85-ce298a79
    api_key = 4fb4441e8388f567cd4a367dae83efa9-69a6bd85-ce298a79
    my domain = sandbox6c87a23cb52845a989d24801f83c6ecb.mailgun.org
    gun domain = https://api.mailgun.net/v3/sandbox6c87a23cb52845a989d24801f83c6ecb.mailgun.org

 source = https://dev-gang.ru/article/izuczaem-go---otpravka-rest-zaprosov-f2b81brpi9/

 recipients
     sosljuk8@gmail.com
     aowler@mail.ru
     cityowler@yandex.ru




curl -s --user 'api:4fb4441e8388f567cd4a367dae83efa9-69a6bd85-ce298a79' \
	https://api.mailgun.net/v3/sandbox6c87a23cb52845a989d24801f83c6ecb.mailgun.org/messages \
	-F from='Excited User <mailgun@YOUR_DOMAIN_NAME>' \
	-F to=sosljuk8@gmail.com \
	-F subject='Hello' \
	-F text='Testing some Mailgun awesomeness!'

	///////////////////////////////////////////////////////////////////////////////////////////

	import (
    	"context"
    	_ "github.com/joho/godotenv"
    	"github.com/mailgun/mailgun-go/v3"
    	"time"
    )

    func main() {

    	// Load environment variables from .env file
    	err := godotenv.Load()

    	if err != nil {
    		log.Fatal("Error loading .env file:", err)
    	}

    	res, erro := SendSimpleMessage(os.Getenv("MG_DOMAIN"), os.Getenv("MG_API_KEY"))

    	fmt.Println(res, erro)

    }

    func SendSimpleMessage(domain, apiKey string) (string, error) {
    	mg := mailgun.NewMailgun(domain, apiKey)
    	m := mg.NewMessage(
    		"Excited User <mailgun@YOUR_DOMAIN_NAME>",
    		"Hello",
    		"Testing some Mailgun awesomeness!bbbbbbbbbbbbbbbbbbbbbbbbFFFFFFFFFFFFFFFFFFFFFFFFFFFFFf",
    		"sosljuk8@gmail.com",
    	)

    	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
    	defer cancel()

    	_, id, err := mg.Send(ctx, m)
    	return id, err
    }

    ////////////////////////////////////////////////////////////