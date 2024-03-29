import { Component, OnInit } from '@angular/core';
import { DashboardService } from '../services/dashboard/dashboard.service';

@Component({
  selector: 'app-portfolio',
  templateUrl: './portfolio.component.html',
  styleUrls: ['./portfolio.component.scss']
})
export class PortfolioComponent implements OnInit {

  internationalList: any;
  displayedColumns: string[] = ['Coin', 'Current International Price', 'Current Indian Price', 'Purchased Quantity', 'Purchase Price', 'Total Amount Spent', 'Average/Coin', 'Change %'];
  purchases: any;
  filteredPortfolio: any;
  wazirList: any;

  constructor(
    private dashboardService: DashboardService
  ) {
    this.dashboardService.getPurchaseInfo("shrujan");
  }

  ngOnInit(): void {
    this.dashboardService.wazirxData$.subscribe(list => {
      if (list) {
        this.wazirList = list;
        this.processPurchasedInfo();
      }
    })
    this.dashboardService.purchaseInfo$.subscribe(purchases => {
      if (purchases) {
        this.purchases = purchases;
        this.processPurchasedInfo();
      }
    })
    this.dashboardService.internationalListInr$.subscribe(list => {
      if (list) {
        this.internationalList = list;
        this.processPurchasedInfo();
      }
    })
  }

  processPurchasedInfo() {
    if (this.wazirList && this.purchases && this.internationalList) {
      this.filteredPortfolio = this.dashboardService.processPurchasedInfo(this.purchases, this.wazirList, this.internationalList);
    }
  }

  calculatePercentIncreaseDecrease(coinPurchasePrice , coinCurrentPrice) {
    coinPurchasePrice = parseInt(coinPurchasePrice);
    coinCurrentPrice  = parseInt(coinCurrentPrice);
    return ((coinCurrentPrice - coinPurchasePrice) / coinPurchasePrice) * 100;
  }

}
