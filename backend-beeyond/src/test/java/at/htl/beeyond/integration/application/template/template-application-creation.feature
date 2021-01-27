Feature: Template application creation endpoint

  Background:
    * url baseUrl
    * path 'application/template'
    * print karate.tags
    * def insertTemplate = read('template-preparation.feature')
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
      "templateId": #(template.id)
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
      "templateId": #(template.id)
    }
    """
    When method POST
    Then status 204

  @student
  Scenario: Create a template application with a too long note
    * def longNote = call read('/at/htl/beeyond/integration:string-generator.js')
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
      "templateId": #(template.id)
    }
    """
    When method POST
    Then status 422
    #* print response
    And match response contains {message: "This field needs to be between 0 and 255 characters",key: "note", value: "#ignore"}
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
      ]
    }
    """
    When method POST
    Then status 422
    And match response == [{message:"This field cannot be empty","value":"",key:"templateId"}]

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
      "templateId": 100
    }
    """
    When method POST
    Then status 422
    #* print response
    And match response == [{message: "Template with id 100 does not exist",value: "100",key: "templateId"}]

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
      "templateId": #(template.id)
    }
    """
    When method POST
    Then status 422
    #* print response
    And match response == [{"message": "Missing field ids: [2], obsolete field ids: []","value": "","key": ""}]

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
        }
      ],
      "templateId": #(template.id)
    }
    """
    When method POST
    Then status 422
    #* print response
    And match response == [{"message": "Missing field ids: [2], obsolete field ids: [3]","value": "","key": ""}]

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
        {
          "fieldId": #(template.fields[1].id)
        }
      ],
      "templateId": #(template.id)
    }
    """
    When method POST
    Then status 422
    #* print response
    And match response contains {"message": "Missing field ids: [2], obsolete field ids: [3]","value": "","key": ""}
    And match response contains {"message": "This field cannot be empty","value": "","key": "value"}

  @student
  Scenario: Create a template with a deleted template
    Given path 'template', template.id
    And request
    And method DELETE
    Then status 200
    Given request
    """
    {
      "note": "string",
      "fieldValues": [
        {
          "fieldId": 1,
          "id": 1,
          "value": "4"
        },
        {
          "fieldId": 2,
          "id": 2,
          "value": "8081"
        }
      ],
      "templateId": 2
    }
    """
    When method POST
    Then status 404

