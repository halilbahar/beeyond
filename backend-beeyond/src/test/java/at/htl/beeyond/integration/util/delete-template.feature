Feature:

  Background:
    * url baseUrl
    * path 'template'
    * configure headers = {Authorization: '#(auth(karate.tags))'}

  @teacher
  Scenario: Delete a valid template
    Given path __arg.id
    When method DELETE
    Then status 204
