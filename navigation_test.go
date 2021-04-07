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

func estoyEnLaPginaAboutUs() error {
	return godog.ErrPending
}

func hagoClickEnElMenu() error {
	divMenu, err := Driver.FindElement(selenium.ByCSSSelector, "#gatsby-focus-wrapper > div > div > header > div > button")
	if err != nil {
		log.Println("Error al traer el menu -> ", err.Error())
	}
	divMenu.Click()
	Driver.SetImplicitWaitTimeout(time.Second * 3)
	return nil
}

func hagoClickEnElEnlanceAboutUs() error {
	aboutUsPath := `//*[@id="gatsby-focus-wrapper"]/div/div/div/div/ul/li[1]/a`
	aboutUs5, err := Driver.FindElement(selenium.ByXPATH, aboutUsPath)

	if err != nil {
		log.Println("Error al hacer click en About Us", err.Error())
	}
	Driver.SetImplicitWaitTimeout(time.Second * 6)
	aboutUs5.Click()

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

	// About Us
	ctx.Step(`^accedo a la pagina principal$`, accedoALaPaginaPrincipal)
	ctx.Step(`^hago click en el menu$`, hagoClickEnElMenu)
	ctx.Step(`^hago click en el enlance About Us$`, hagoClickEnElEnlanceAboutUs)
	ctx.Step(`^estoy en la página About Us$`, estoyEnLaPginaAboutUs)
	ctx.Step(`^estoy en la página About Us con texto "([^"]*)"$`, estoyEnLaPginaAboutUsConTexto)

	// Invalid Email
	ctx.Step(`^ingreso un email con formato incorrecto "([^"]*)"$`, ingresoUnEmailConFormatoIncorrecto)
	ctx.Step(`^al hacer click en el botón \'Suscribe\'$`, alHacerClickEnElBotnSuscribe)
	ctx.Step(`^aparece un mensaje "([^"]*)"$`, apareceUnMensaje)

}
