package main

import (
	"fmt"
	"log"
	"time"

	"github.com/cucumber/godog"
	"github.com/jonathanbs9/go-selenium-avalith-web-test/support"
	"github.com/tebeka/selenium"
)

var (
	Driver selenium.WebDriver
)

func accedoALaPaginaPrincipal() error {
	log.Println("Acceso a página avalith.net ")
	Driver.Get("https://avalith.net/")

	return nil
}

func hagoClickEnElMenu() error {
	divMenu, err := Driver.FindElement(selenium.ByCSSSelector, "#gatsby-focus-wrapper > div > div > header > div > button")
	if err != nil {
		log.Println("Error al traer el menu -> ", err.Error())
	}
	divMenu.Click()
	Driver.SetImplicitWaitTimeout(time.Second * 6)
	return nil
}

func hagoClickEnElEnlanceAboutUs() error {
	aboutUs5, err := Driver.FindElement(selenium.ByCSSSelector, `#gatsby-focus-wrapper > div > div > div > div > ul > li:nth-child(1) > a`)

	if err != nil {
		log.Println("Error al hacer click en About Us", err.Error())
	}
	Driver.SetImplicitWaitTimeout(time.Second * 10)
	aboutUs5.Click()

	return nil
}

func hagoClickEnElEnlaceServices() error {
	serviceXPath := `//*[@id="gatsby-focus-wrapper"]/div/div/div/div/ul/li[2]/a`
	service, err := Driver.FindElement(selenium.ByXPATH, serviceXPath)

	if err != nil {
		log.Println("Error al hacer click en Service", err.Error())
	}
	Driver.SetImplicitWaitTimeout(time.Second * 10)
	service.Click()

	return nil
}

func hagoClickEnElEnlaceCareers() error {
	careerXPATH := `//*[@id="gatsby-focus-wrapper"]/div/div/div/div/ul/li[3]/a`
	career, err := Driver.FindElement(selenium.ByXPATH, careerXPATH)
	if err != nil {
		log.Println("Error al hacer click en Career", err.Error())
	}
	Driver.SetImplicitWaitTimeout(time.Second * 10)
	career.Click()
	return nil
}

func estoyEnLaPginaAboutUsConTexto(text string) error {
	aboutUsdiv, err := Driver.FindElement(selenium.ByCSSSelector, `#gatsby-focus-wrapper > div > main > div.container.mx-auto.md\:mt-20.px-6 > section.flex.flex-col-reverse.lg\:flex-row.items-center > div:nth-child(1) > h3`)
	if err != nil {
		log.Println(err)
	}
	aboutUsText, _ := aboutUsdiv.Text()
	if aboutUsText != text {
		log.Fatal("No son iguales los textos")
		log.Fatalf("mensaje obtenido : %s  | Mensaje esperado: %s", aboutUsText, text)
	}

	return nil
}

func estoyEnLaPginaServicesConTexto(textServicePage string) error {
	servicesdiv, err := Driver.FindElement(selenium.ByCSSSelector, `#gatsby-focus-wrapper > div > main > section.container.mx-auto.px-6.py-8 > div.flex.flex-col.md\:flex-row.items-center.pb-8 > div.w-full.md\:w-1\/2.md\:px-8 > p`)
	if err != nil {
		log.Println(err)
	}
	servicesText, _ := servicesdiv.Text()
	if servicesText != textServicePage {
		log.Fatalf("mensaje obtenido : %s  | Mensaje esperado: %s", textServicePage, servicesText)
	}

	return nil
}

func estoyEnLaPginaCareersConTexto(careerText string) error {
	careerElement, err := Driver.FindElement(selenium.ByCSSSelector, `#gatsby-focus-wrapper > div > main > section.container.mx-auto.md\:mt-20.px-6.pb-20 > div.flex.flex-col.md\:flex-row.items-center.pb-20 > div.w-full.md\:w-1\/2.md\:px-8 > h5`)
	if err != nil {
		log.Println(err)
	}
	careerElementText, _ := careerElement.Text()
	if careerElementText != careerText {
		log.Fatalf("mensaje obtenido : %s  | Mensaje esperado: %s", careerElementText, careerText)
	}

	return nil
}

func ingresoUnEmailConFormatoIncorrecto(invalidEmail string) error {
	campoEmail, err := Driver.FindElement(selenium.ByXPATH, `//*[@id="gatsby-focus-wrapper"]/div/footer/div[1]/div[2]/form/input`)
	if err != nil {
		log.Println("Error al encontrar input ")
	}
	Driver.SetImplicitWaitTimeout(time.Second * 5)
	campoEmail.SendKeys(invalidEmail)

	return nil
}

func alHacerClickEnElBotnSuscribe() error {
	buttonSubscribe, err := Driver.FindElement(selenium.ByXPATH, `//*[@id="gatsby-focus-wrapper"]/div/footer/div[1]/div[2]/form/button`)
	if err != nil {
		log.Println("Error al buscar botón de suscripción")
	}
	buttonSubscribe.Click()

	return nil
}

func apareceUnMensaje(msg string) error {
	messageAlert, err := Driver.FindElement(selenium.ByXPATH, `//*[@id="gatsby-focus-wrapper"]/div/footer/div[1]/div[2]/p`)
	if err != nil {
		log.Println("Error al obtener mensaje de alerta")
	}

	alerta, err := messageAlert.Text()
	if err != nil {
		log.Println("Error al recibir el texto del alerta")
	}

	if alerta != msg {
		return fmt.Errorf("El mensaje obtenido no es el mensaje esperado")
	}

	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.BeforeScenario(func(*godog.Scenario) {
		Driver = support.WDInit()
	})

	ctx.Step(`^accedo a la pagina principal$`, accedoALaPaginaPrincipal)
	ctx.Step(`^hago click en el menu$`, hagoClickEnElMenu)

	// About Us - Services
	ctx.Step(`^hago click en el enlance About Us$`, hagoClickEnElEnlanceAboutUs)
	ctx.Step(`^hago click en el enlace Services$`, hagoClickEnElEnlaceServices)
	ctx.Step(`^hago click en el enlace careers$`, hagoClickEnElEnlaceCareers)

	ctx.Step(`^estoy en la página About Us con texto "([^"]*)"$`, estoyEnLaPginaAboutUsConTexto)
	ctx.Step(`^estoy en la página Services con texto "([^"]*)"$`, estoyEnLaPginaServicesConTexto)
	ctx.Step(`^estoy en la página careers con texto "([^"]*)"$`, estoyEnLaPginaCareersConTexto)

	// Invalid Email
	ctx.Step(`^ingreso un email con formato incorrecto "([^"]*)"$`, ingresoUnEmailConFormatoIncorrecto)
	ctx.Step(`^al hacer click en el botón \'Suscribe\'$`, alHacerClickEnElBotnSuscribe)
	ctx.Step(`^aparece un mensaje "([^"]*)"$`, apareceUnMensaje)

	ctx.AfterScenario(func(sc *godog.Scenario, err error) {
		log.Println("Closing Chrome!")
		Driver.Quit()
	})
}
