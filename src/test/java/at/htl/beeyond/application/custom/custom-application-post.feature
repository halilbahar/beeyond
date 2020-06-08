Feature: Test for creating a custom application

  Background:
    Given url baseUrl
    Given path '/authentication/login'
    And request { username: 'teststudent', password: 'student'}
    When method post
    * def accessToken = response.access_token

  @validCustomApplication
  Scenario: Create a custom application
    Given path '/application/custom'
    And header Authorization = 'Bearer ' + accessToken
    And header Accept = 'application/json'
    And request {name: 'Nginx 1.16',description: 'This is a simple template for a nginx server', content: 'hello %port% %name%', fields: [{wildcard: 'port',label: 'Port of the webserver',description: 'Lorem ipsum dolor sit amet, consectetuer adipiscing elit.'},{wildcard: 'name',label: 'Name of my server',description: 'pretium quis'}], note: 'This is a note'}
    When method post
    Then status 204

  @blankContent
  Scenario: Create an custom application with blank content
    Given path '/application/custom'
    And header Authorization = 'Bearer ' + accessToken
    And header Content-Type = 'application/json'
    And request {name: 'Nginx 1.16',description: 'This is a simple template for a nginx server', content: '', fields: [{wildcard: 'port',label: 'Port of the webserver',description: 'Lorem ipsum dolor sit amet, consectetuer adipiscing elit.'},{wildcard: 'name',label: 'Name of my server',description: 'pretium quis'}], note: 'This is a note'}
    When method post
    Then status 422