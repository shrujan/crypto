import { Component, OnInit } from '@angular/core';
import { DashboardService } from '../services/dashboard/dashboard.service';

@Component({
  selector: 'app-favorites',
  templateUrl: './favorites.component.html',
  styleUrls: ['./favorites.component.scss']
})
export class FavoritesComponent implements OnInit {

  internationalList: any;
  displayedColumns: string[] = ['Coin', 'Current International Price', 'Current Indian Price'];
  favCoins: any;
  wazirList: any;

  constructor(
    private dashboardService: DashboardService
  ) {}

  ngOnInit(): void {
    this.dashboardService.getCoins();
    this.dashboardService.wazirxData$.subscribe(list => {
      if (list) {
        this.wazirList = list;
        console.log('wazirList = ',this.wazirList)
      }
    })

    this.dashboardService.internationalListInr$.subscribe(list => {
      if (list) {
        this.internationalList = list;
        console.log('internationsl = ',this.internationalList)
      }
    })

    this.dashboardService.favoriteCoins$.subscribe(favCoinsList => {
      if (favCoinsList) {
        this.favCoins = favCoinsList;
        console.log(this.favCoins)
      }
    })
  }

  getInternationalCoinPrice(coin) {
    return (coin && this.internationalList) ? this.internationalList.find(internationlCoin => internationlCoin.symbol === coin.Symbol).current_price : '';
  }

  getWXCoinPrice(coin) {
    return (coin && this.wazirList) ? this.wazirList[coin.Symbol+"inr"].buy : '';
  }

}
