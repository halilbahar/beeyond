Feature: Create a template application

  Background:
    Given url baseUrl
    Given path '/authentication/login'
    And request { username: 'teststudent', password: 'student'}
    When method post
    * def accessToken = response.access_token
  
  Scenario: Create a template application
    Given path '/application/template'
    And header Authorization = 'Bearer ' + accessToken
    And header Content-Type = 'application/json'
    And request {templateId: 1, name: 'Nginx 1.16',description: 'This is a simple template for a nginx server', content: 'Some content', fields: [{wildcard: 'port',label: 'Port of the webserver',description: 'Lorem ipsum dolor sit amet, consectetuer adipiscing elit.'},{wildcard: 'name',label: 'Name of my server',description: 'pretium quis'}], note: 'Some note'}
    When method post
    Then status 204


  