import { Component, OnInit } from '@angular/core';
import { DashboardService } from '../services/dashboard/dashboard.service';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.scss']
})
export class DashboardComponent implements OnInit {
  coins: any;
  filteredCoins: any;

  constructor(private dashboardService: DashboardService) { }

  ngOnInit(): void {
    this.getCoinData();
  }

  getCoinData() {
    console.log('coins')
    this.dashboardService.getCoinsInfoInr().subscribe((data) => {
      this.coins = data;
      this.filteredCoins = data;
    });
  }
}
