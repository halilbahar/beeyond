import { Component, OnInit } from '@angular/core';
import { FormControl, Validators } from '@angular/forms';

@Component({
  selector: 'app-application-deny-dialog',
  templateUrl: './application-deny-dialog.component.html',
  styleUrls: ['./application-deny-dialog.component.scss']
})
export class ApplicationDenyDialogComponent implements OnInit {
  message = new FormControl('', Validators.required);

  constructor() {}

  ngOnInit(): void {}
}
