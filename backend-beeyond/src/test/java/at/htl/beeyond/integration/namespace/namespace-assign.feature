Feature: Namespace assigning endpoint

  Background:
    * url baseUrl + '/namespace'
    * configure headers = { Authorization: '#(auth(karate.tags))' }

  @teacher
  Scenario: Assign 1 user to a valid namespace
    Given request
    """
    {
      "namespace": "beeyond",
      "users": [
        "emina"
      ]
    }
    """
    When method PUT
    Then status 204
    And path 'beeyond'
    When method GET
    Then status 200
    And match response.namespace == 'beeyond'
    And match response.users == '#[1]'

  @teacher
  Scenario: Assign 2 user to a valid namespace
    Given request
    """
    {
      "namespace": "beeyond",
      "users": [
        "emina",
        "moritz"
      ]
    }
    """
    When method PUT
    Then status 204
    And path 'beeyond'
    When method GET
    Then status 200
    And match response.namespace == 'beeyond'
    And match response.users == '#[2]'

  @teacher
  Scenario: Assign 3 user to a valid namespace
    Given request
    """
    {
      "namespace": "beeyond",
      "users": [
        "emina",
        "moritz",
        "marc"
      ]
    }
    """
    When method PUT
    Then status 204
    And path 'beeyond'
    When method GET
    Then status 200
    And match response.namespace == 'beeyond'
    And match response.users == '#[3]'

  @teacher
  Scenario: Assign 2 user to a invalid namespace (wrong characters)
    Given request
    """
    {
      "namespace": "bééyond",
      "users": [
        "emina",
        "moritz"
      ]
    }
    """
    When method PUT
    Then status 422

#  TODO: better string generator
#  @teacher
#  Scenario: Assign 2 user to a invalid namespace (too long)
#    * def generateString = read('string-generator.js')
#    Given request
#    """
#    {
#      "namespace": #(generateString()),
#      "users": [
#        "emina",
#        "moritz"
#      ]
#    }
#    """
#    When method PUT
#    Then status 422
#    And match response[0].message == 'This field needs to be between 1 and 253 characters'

  @teacher
  Scenario: Assign 2 user to a invalid namespace (name of user)
    Given request
    """
    {
      "namespace": "marc",
      "users": [
        "emina",
        "moritz"
      ]
    }
    """
    When method PUT
    Then status 422
    And match response[0].message == 'Not a valid namespace name'

  @teacher
  Scenario: Assign a nonexistent user to a valid namespace
    Given request
    """
    {
      "namespace": "beeyond",
      "users": [
         "test"
      ]
    }
    """
    When method PUT
    Then status 422
    And match response[0].message == 'User with name [test] does not exist'

  @teacher
  Scenario: Assign a user to its own namespace (name of user)
    Given request
    """
    {
      "namespace": "emina",
      "users": [
        "emina"
      ]
    }
    """
    When method PUT
    Then status 422
    And match response[0].message == 'Not a valid namespace name'

  @teacher
  Scenario: Assign 2 user to a valid namespace where one of them already belongs to that namespace
    Given request
    """
    {
      "namespace": "beeyond",
      "users": [
        "emina"
      ]
    }
    """
    When method PUT
    Given request
    """
    {
      "namespace": "beeyond",
      "users": [
        "emina",
        "marc"
      ]
    }
    """
    When method PUT
    And path 'beeyond'
    When method GET
    Then status 200
    And match response.namespace == 'beeyond'
    And match response.users == '#[2]'

  @teacher
  Scenario: Assign 2 user to a valid namespace where both of them already belong to that namespace
    Given request
    """
    {
      "namespace": "beeyond",
      "users": [
        "emina",
        "marc"
      ]
    }
    """
    When method PUT
    Given request
    """
    {
      "namespace": "beeyond",
      "users": [
        "emina",
        "marc"
      ]
    }
    """
    When method PUT
    And path 'beeyond'
    When method GET
    Then status 200
    And match response.namespace == 'beeyond'
    And match response.users == '#[2]'

  @teacher
  Scenario: Assign 2 user with the same name (same user) to a valid namespace
    Given request
    """
    {
      "namespace": "beeyond",
      "users": [
        "emina",
        "emina"
      ]
    }
    """
    When method PUT
    And path 'beeyond'
    When method GET
    Then status 200
    And match response.namespace == 'beeyond'
    And match response.users == '#[1]'

  @teacher
  Scenario: Assign 2 user with the same name (same user) to a invalid namespace (name of user)
    Given request
    """
    {
      "namespace": "emina",
      "users": [
        "emina",
        "emina"
      ]
    }
    """
    When method PUT
    Then status 422
    And match response[0].message == 'Not a valid namespace name'

  @teacher
  Scenario: Assign 1 user and remove 1 from a valid namespace
    Given request
    """
    {
      "namespace": "beeyond",
      "users": [
        "emina"
      ]
    }
    """
    When method PUT
    Then status 204
    Given request
    """
    {
      "namespace": "beeyond",
      "users": [
        "marc"
      ]
    }
    """
    When method PUT
    Then status 204
    And path 'beeyond'
    When method GET
    Then status 200
    And match response.namespace == 'beeyond'
    And match response.users == '#[1]'
    And match response.users[0].name == 'marc'

  @teacher
  Scenario: Remove 1 from a valid namespace (1 user exists already)
    Given request
    """
    {
      "namespace": "beeyond",
      "users": [
        "emina"
      ]
    }
    """
    When method PUT
    Then status 204
    Given request
    """
    {
      "namespace": "beeyond",
      "users": []
    }
    """
    When method PUT
    Then status 204
    And path 'beeyond'
    When method GET
    Then status 200
    And response.deleted == true

  @teacher
  Scenario: Assign 1 to a valid namespace (1 user exists already)
    Given request
    """
    {
      "namespace": "beeyond",
      "users": [
        "emina"
      ]
    }
    """
    When method PUT
    Then status 204
    Given request
    """
    {
      "namespace": "beeyond",
      "users": [
        "emina",
        "marc"
      ]
    }
    """
    When method PUT
    Then status 204
    And path 'beeyond'
    When method GET
    Then status 200
    And match response.namespace == 'beeyond'
    And match response.users == '#[2]'

  @teacher
  Scenario: Replace 2 users with 2 other users (A, B -> C, D)
    Given request
    """
    {
      "namespace": "beeyond",
      "users": [
        "emina",
        "marc"
      ]
    }
    """
    When method PUT
    Then status 204
    Given request
    """
    {
      "namespace": "beeyond",
      "users": [
        "moritz",
        "stuetz"
      ]
    }
    """
    When method PUT
    Then status 204
    And path 'beeyond'
    When method GET
    Then status 200
    And match response.namespace == 'beeyond'
    And match response.users == '#[2]'
    And match response.users contains { 'name': 'moritz', 'id': #number }
    And match response.users contains { 'name': 'stuetz', 'id': #number }

  Scenario: Create a namespace with empty user
