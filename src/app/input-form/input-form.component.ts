import { Component, OnInit } from '@angular/core';
import { DashboardService, User } from '../services/dashboard/dashboard.service';

@Component({
  selector: 'app-input-form',
  templateUrl: './input-form.component.html',
  styleUrls: ['./input-form.component.scss']
})
export class InputFormComponent implements OnInit {
  userList: any;
  selected: string = '';
  constructor(
    private dashboardService: DashboardService
  ) {
      this.dashboardService.getUsers().subscribe(userList => {
        this.userList = userList;
      })
   }

  ngOnInit(): void {
  }

  saveDetails(): void {

  }

}
