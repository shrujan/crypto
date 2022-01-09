import { Component } from '@angular/core';
import { DashboardService } from './services/dashboard/dashboard.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent {
  title = 'crypto';
  
  constructor(
    private dashboardService: DashboardService
  ) {
    this.dashboardService.getInternationalData();
    this.dashboardService.fetchAPIWazirx();
    this.dashboardService.getCoins();
  }
}
