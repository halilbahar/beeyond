Feature:

  Background:
    * url baseUrl
    * configure headers = {Authorization: '#(auth(karate.tags))'}

  @student
  Scenario: Stop application
    Given path 'application/stop/'+id
    When method PATCH
    Then status 200
