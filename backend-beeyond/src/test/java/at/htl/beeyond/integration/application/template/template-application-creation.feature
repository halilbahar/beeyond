Feature: Template application creation endpoint

  Background:
    * url baseUrl
    * path 'application/template'
    * configure headers = { Authorization: '#(auth(karate.tags))' }
    * def insertTemplate = read('classpath:at/htl/beeyond/integration/util/create-template.feature')
    * def insertTemplateResponse = call insertTemplate
    * def template = insertTemplateResponse.template

  @student
  Scenario: Create a valid template application
    Given request
    """
    {
      "note": "string",
      "fieldValues": [
        {
          "fieldId": #(template.fields[0].id),
          "value": "4"
        },
        {
          "fieldId": #(template.fields[1].id),
          "value": "8081"
        }
      ],
      "templateId": #(template.id),
      "namespace": "moritz"
    }
    """
    When method POST
    Then status 204

  @student
  Scenario: Create a template application with no note
    Given request
    """
    {
      "fieldValues": [
        {
          "fieldId": #(template.fields[0].id),
          "value": "4"
        },
        {
          "fieldId": #(template.fields[1].id),
          "value": "8081"
        }
      ],
      "templateId": #(template.id),
      "namespace": "moritz"
    }
    """
    When method POST
    Then status 204

  @student
  Scenario: Create a template application with a too long note
    * def longNote = call read('classpath:string-generator.js')
    Given request
    """
    {
      "note": "#(longNote)",
      "fieldValues": [
        {
          "fieldId": #(template.fields[0].id),
          "value": "4"
        },
        {
          "fieldId": #(template.fields[1].id),
          "value": "8081"
        }
      ],
      "templateId": #(template.id),
      "namespace": "moritz"
    }
    """
    When method POST
    Then status 422
    And match response contains
    """
    {
      message: "This field needs to be between 0 and 255 characters",
      key: "note",
      value: "#ignore"
    }
    """

  @student
  Scenario: Create a template application with no templateId
    Given request
    """
    {
      "note": "string",
      "fieldValues": [
        {
          "fieldId": #(template.fields[0].id),
          "value": "4"
        },
        {
          "fieldId": #(template.fields[1].id),
          "value": "8081"
        }
      ],
      "namespace": "moritz"
    }
    """
    When method POST
    Then status 422
    And match response contains
    """
    {
      "message": "This field cannot be empty",
      "value": "",
      "key": "templateId"
    }
    """

  @student
  Scenario: Create a template application with a non existing templateId
    Given request
    """
    {
      "note": "string",
      "fieldValues": [
        {
          "fieldId": #(template.fields[0].id),
          "value": "4"
        },
        {
          "fieldId": #(template.fields[1].id),
          "value": "8081"
        }
      ],
      "templateId": 10000,
      "namespace": "moritz"
    }
    """
    When method POST
    Then status 422
    And match response contains
    """
    {
      message: "Template with id 10000 does not exist",
      value: "10000",
      key: "templateId"
    }
    """

  @student
  Scenario: Create a template application where at least 1 field is missing
    Given request
    """
    {
      "note": "string",
      "fieldValues": [
        {
          "fieldId": #(template.fields[0].id),
          "value": "4"
        }
      ],
      "templateId": #(template.id),
      "namespace": "moritz"
    }
    """
    When method POST
    Then status 422
    And def message = 'Missing field ids: [' + template.fields[1].id + '], obsolete field ids: []'
    And def expected =
    """
    {
      message: "#(message)",
      value: "",
      key: ""
    }
    """
    And match response contains expected

  @student
  Scenario: Create a template application where all fields are provided but one with an invalid fieldId
    Given request
    """
    {
      "note": "string",
      "fieldValues": [
        {
          "fieldId": #(template.fields[0].id),
          "value": "4"
        },
        {
          "fieldId": 9999,
          "value": "8081"
        },
        {
          "fieldId": 10000,
          "value": 8082
        },
        {
          "fieldId": 10001,
          "value": 8083
        }
      ],
      "templateId": #(template.id),
      "namespace": "moritz"
    }
    """
    When method POST
    Then status 422
    And def message = 'Missing field ids: [' + template.fields[1].id + '], obsolete field ids: [9999, 10000, 10001]'
    And def expected =
    """
    {
      message: "#(message)",
      value: "",
      key: ""
    }
    """
    And match response contains expected

  @student
  Scenario: Create a template application where at least 1 field is provided with no value and fieldId
    Given request
    """
    {
      "note": "string",
      "fieldValues": [
        {
          "fieldId": #(template.fields[0].id),
          "value": "4"
        },
        {}
      ],
      "templateId": #(template.id),
      "namespace": "moritz"
    }
    """
    When method POST
    Then status 422
    And match response contains { "message": "This field cannot be empty", "value": "", "key": "value" }

  @student
  Scenario: Create a template application where at least 1 field is provided with no fieldId
    Given request
    """
    {
      "note": "string",
      "fieldValues": [
        {
          "fieldId": #(template.fields[0].id),
          "value": "4"
        },
        {
          "value": 8080
        }
      ],
      "templateId": #(template.id),
      "namespace": "moritz"
    }
    """
    When method POST
    Then status 422
    And match response contains { "message": "This field cannot be empty", "value": "", "key": "fieldId" }

  @student
  Scenario: Create a valid template application with a deleted template
    * def deleteTemplate = read('classpath:at/htl/beeyond/integration/util/delete-template.feature')
    * def insertTemplateResponse = call deleteTemplate { id: #(template.id) }
    Given path 'application', 'template'
    Given request
    """
    {
      "note": "string",
      "fieldValues": [
        {
          "fieldId": #(template.fields[0].id),
          "value": "4"
        },
        {
          "fieldId": #(template.fields[0].id),
          "value": "8081"
        }
      ],
      "templateId": #(template.id),
      "namespace": "moritz"
    }
    """
    When method POST
    Then status 404

  @student
  Scenario: Create a invalid template application with a deleted template where the fieldValues are missing
    * def deleteTemplate = read('classpath:at/htl/beeyond/integration/util/delete-template.feature')
    * def insertTemplateResponse = call deleteTemplate { id: #(template.id) }
    Given path 'application', 'template'
    Given request
    """
    {
      "note": "string",,
      "namespace": "moritz"
      "templateId": #(template.id)
    }
    """
    When method POST
    Then status 404

  @student
  Scenario: Create a template application where the id of a template field is set
