Feature: Custom application creation endpoint

  Background:
    * url baseUrl
    * path 'application/custom'
    * def nginxDeployment = read('classpath:at/htl/beeyond/integration/util/nginx-deployment.yml.txt')
    * configure headers = {Authorization: '#(auth(karate.tags))'}

  @student
  Scenario: Create a valid custom application
    Given request
    """
    {
      "note": "Nginx Deployment",
      "content": "#(nginxDeployment)",
      "namespace": "moritz",
      "class": "5AHIF",
      "to": "18.01.2022",
      "purpose": "SYP"
    }
    """
    When method POST
    Then status 201

  @student
  Scenario: Create a custom application with no note
    Given request
    """
    {
      "content": "#(nginxDeployment)",
      "namespace": "moritz",
      "class": "5AHIF",
      "to": "18.01.2022",
      "purpose": "SYP"
    }
    """
    When method POST
    Then status 201

  @student
  Scenario: Create a custom application with no content
    Given request
    """
    {
      "namespace": "moritz",
      "class": "5AHIF",
      "to": "18.01.2022",
      "purpose": "SYP"
    }
    """
    When method POST
    Then status 422
    And match response contains { "message": "This field cannot be empty", "value": "", "key": "content" }

  @student
  Scenario: Create a custom application with blank content
    Given request
    """
    {
      "content": "",
      "namespace": "moritz",
      "class": "5AHIF",
      "to": "18.01.2022",
      "purpose": "SYP"
    }
    """
    When method POST
    Then status 422
    And match response contains { "message": "This field cannot be empty", "value": "", "key": "content" }

  @student
  Scenario: Create a custom application with invalid kubernetes content

  @student
  Scenario: Create a custom application with a too long note
    Given request
    """
    {
      "note": "#(generateString())",
      "content": "#(content)",
      "namespace": "moritz",
      "class": "5AHIF",
      "to": "18.01.2022",
      "purpose": "SYP"
    }
    """
    When method POST
    Then status 422
    And match response contains { "message": "This field needs to be between 0 and 255 characters", "value": "#ignore", "key":"note"}

  @student
  Scenario: Create a custom application with id set
    Given request
    """
    {
      "id": 9999,
      "note": "Nginx Deployment",
      "content": "#(nginxDeployment)",
      "namespace": "moritz",
      "class": "5AHIF",
      "to": "18.01.2022",
      "purpose": "SYP"
    }
    """
    When method POST
    Then status 201

  @student
  Scenario: Create a custom application with application status set
    Given request
    """
    {
      "status": "APPROVED",
      "note": "Nginx Deployment",
      "content": "#(nginxDeployment)",
      "namespace": "moritz",
      "class": "5AHIF",
      "to": "18.01.2022",
      "purpose": "SYP"
    }
    """
    When method POST
    Then status 201

  @student
  Scenario: Create a custom application with owner set
    Given request
    """
    {
      "owner": {
        "name": "stuetz",
        "id": 9
      },
      "note": "Nginx Deployment",
      "content": "#(nginxDeployment)",
      "namespace": "moritz",
      "class": "5AHIF",
      "to": "18.01.2022",
      "purpose": "SYP"
    }
    """
    When method POST
    Then status 201

  @student
  Scenario: Create a custom application with created at set
    Given request
    """
    {
      "createdAt": "2021-02-10T22:42:57.620598",
      "note": "Nginx Deployment",
      "content": "#(nginxDeployment)",
      "namespace": "moritz",
      "class": "5AHIF",
      "to": "18.01.2022",
      "purpose": "SYP"
    }
    """
    When method POST
    Then status 201
