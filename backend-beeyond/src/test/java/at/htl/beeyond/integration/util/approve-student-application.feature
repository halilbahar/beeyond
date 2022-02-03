Feature:

  Background:
    * url baseUrl
    * configure headers = {Authorization: '#(auth(karate.tags))'}

  @teacher
  Scenario: Approve application
    Given path 'application/approve/'+id
    When method PATCH
    Then status 200