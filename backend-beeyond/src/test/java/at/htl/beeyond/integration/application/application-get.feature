Feature: Get application endpoint

  Background:
    * url baseUrl
    * path 'application'
    * def insertApplication = read('classpath:at/htl/beeyond/integration/util/create-application.feature')
    * call insertApplication
    * configure headers = {Authorization: '#(auth(karate.tags))'}

  @teacher
  Scenario: Get all applications
    When method GET
    Then status 200
    And assert response.length == 1

  @student
  Scenario: Get all applications as student
    When method GET
    Then status 200
    And assert response.length == 0
