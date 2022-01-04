import { Component, OnInit } from '@angular/core';
import { DashboardService } from '../services/dashboard/dashboard.service';
import { find, findIndex } from 'lodash'; 

interface FavoriteCoin {
  id: string,
  favorite: boolean,
  purchasedPrice: number,
}

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.scss']
})
export class DashboardComponent implements OnInit {
  coins: any;
  displayedColumns: string[] = ['Favorites', 'Coin', 'Price', '24 High', '24 Low', 'Change %', 'Purchased'];
  favorites: FavoriteCoin[] = [];
  filteredCoins: any;

  constructor(private dashboardService: DashboardService) { }

  ngOnInit(): void {
    this.dashboardService.getCoinsInfoInr();
    this.dashboardService.coinListInr$.subscribe(coinList => {
      if (!coinList) return
      this.setCustomPrices(coinList);
      this.coins = coinList;
      this.filteredCoins = coinList;
    })
    // const source = interval(1000000);
    // source.subscribe(() => this.getCoinData())
    
  }

  savePurchasedPrice(coin: any): void {
    const updatePrice = (favCoin: FavoriteCoin) => favCoin.purchasedPrice = coin.purchasedPrice;
    const favCoin = find(this.favorites, fav => fav.id === coin.id);
    (!favCoin) ? this.saveNewFavCoin(coin) : updatePrice(favCoin);
    this.setCustomPrices(this.filteredCoins)
  }

  setCustomPrices(coinList: any): void {
    coinList.forEach((coin: any) => {
      const favCoin = find(this.favorites, fav => fav.id === coin.id)
      coin.favorite = false;
      coin.purchasedPrice = 0;
      coin.differenceFromPurchase = 0;
      if (favCoin) {
        coin.favorite = true;
        coin.purchasedPrice = favCoin.purchasedPrice,
        coin.differenceFromPurchase = this.calcPercentage(Number(coin.purchasedPrice), Number(coin.current_price))
      }
    });
  }

  setFavorite(coin: any) {
    const deleteFavCoin = (index: number) => this.favorites.splice(index, 1);
    const favCoinIndex = findIndex(this.favorites, (fav: FavoriteCoin) => fav.id === coin.id);
    (favCoinIndex === -1) ? this.saveNewFavCoin(coin) : deleteFavCoin(favCoinIndex);
    this.setCustomPrices(this.filteredCoins)
  }

  calcPercentage(purchasedPrice: number, currentPrice: number): number {
    return (!purchasedPrice) ? 0 : (currentPrice - purchasedPrice) / purchasedPrice * 100;
  }

  saveNewFavCoin(coin: any): void {
    this.favorites.push({
      id: coin.id,
      favorite: true,
      purchasedPrice: coin.purchasedPrice,
    });
  }
}
