export class BlueprintTemplate {

  public name = '';
  public replica = 0;
  public templateName = '';

  constructor(name: string, replica: number, templateName: string) {
    this.name = name;
    this.replica = replica;
    this.templateName = templateName;
  }
}
