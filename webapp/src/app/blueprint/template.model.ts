export class Template {

  name = '';
  description = '';
  content = '';

  constructor(name: string, description: string, content: string) {
    this.name = name;
    this.description = description;
    this.content = content;
  }
}
