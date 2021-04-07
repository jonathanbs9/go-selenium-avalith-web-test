Feature: Navegación
    Para que pueda acceder a la pagína
    Siendo un usuario cualquiera

    @AboutUs
    Scenario: Navegacion a About Us
        Given accedo a la pagina principal
        When hago click en el menu
        And hago click en el enlance About Us
        Then estoy en la página About Us con texto "Two brothers with the same passion, the same dream: building their careers and enhancing the IT industry."
    
    @Services
    Scenario: Navegacion a Services
        Given accedo a la pagina principal
        When hago click en el menu
        And hago click en el enlace Services
        Then estoy en la página Services con texto "From Avalith we offer you support at times where you need it most, with complete support for all your projects."

    @Careers
    Scenario: Navegacion a Careers
        Given accedo a la pagina principal
        When hago click en el menu
        And hago click en el enlace careers
        Then estoy en la página careers con texto ""

    @OurPartners
    Scenario: Navegacion a Our Partners
        Given accedo a la pagina principal
        When hago click en el menu
        And hago click en el enlace our partners
        Then estoy en la página Our Partners con texto ""

    @ContactUs
    Scenario: Navegacion a Contact Us
        Given accedo a la pagina principal
        When hago click en el menu
        And hago click en el enlace Contact Us 
        Then estoy en la página Contact Us con texto ""

    @InvalidEmailSubscription
    Scenario: Email inválido
        Given accedo a la pagina principal
        When ingreso un email con formato incorrecto "invalidemail@avalith"
        And al hacer click en el botón 'Suscribe'
        Then aparece un mensaje "The email you entered is not valid."
    