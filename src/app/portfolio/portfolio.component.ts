import { Component, OnInit } from '@angular/core';
import { DashboardService } from '../services/dashboard/dashboard.service';

@Component({
  selector: 'app-portfolio',
  templateUrl: './portfolio.component.html',
  styleUrls: ['./portfolio.component.scss']
})
export class PortfolioComponent implements OnInit {

  coinList: any;
  displayedColumns: string[] = ['Coin', 'Current International Price', 'Current Indian Price', 'Purchased Date', 'Purchased Quantity', 'Purchase Price', 'Total Amount', 'Change %'];
  purchases: any;
  filteredPortfolio: any;

  constructor(
    private dashboardService: DashboardService
  ) {
    
    
  }

  ngOnInit(): void {
    this.dashboardService.fetchAPIWazirx();
    this.dashboardService.getPurchaseInfo("shrujan");
    this.dashboardService.wazirxData$.subscribe(list => {
      if (list) {
        this.coinList = list;
        this.processPurchasedInfo();
      }
    })
    this.dashboardService.purchaseInfo$.subscribe(purchases => {
      if (purchases) {
        this.purchases = purchases;
        this.processPurchasedInfo();
      }
    })
    // this.dashboardService.coinListInr$.subscribe(coinList => {
    //   if (!coinList) return
    //   this.setCustomPrices(coinList);
    //   this.coins = coinList;
    //   this.filteredCoins = coinList;
    // })
  }

  processPurchasedInfo() {
    if (this.coinList && this.purchases) {
      this.filteredPortfolio = this.dashboardService.processPurchasedInfo(this.purchases, this.coinList);
    }
  }

}
