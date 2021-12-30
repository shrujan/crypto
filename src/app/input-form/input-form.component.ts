import { Component } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';
import { DashboardService, User } from '../services/dashboard/dashboard.service';

@Component({
  selector: 'app-input-form',
  templateUrl: './input-form.component.html',
  styleUrls: ['./input-form.component.scss']
})
export class InputFormComponent {
  userList: any;
  selected: string = '';
  cryptoForm: FormGroup;

  constructor(
    private dashboardService: DashboardService,
    private formBuilder: FormBuilder,
  ) {
      this.dashboardService.getUsers().subscribe(userList => {
        this.userList = userList;
      });
      this.initializeForm();
   }

  initializeForm() {
    this.cryptoForm = new FormGroup({
      userName: new FormControl(''),
      coinName: new FormControl(),
      quantity: new FormControl(),
      purchasePrice: new FormControl(),
      purchaseDate: new FormControl(),
      totalAmount: new FormControl()
    })
  }

  saveDetails(): void {
    const cryptoData = this.cryptoForm.value;
    const param = {
      "userName": (cryptoData.userName.UserName).toLowerCase(),
      "coinName": (cryptoData.coinName).toLowerCase(),
      "quantity": cryptoData.quantity,
      "purchasePrice": cryptoData.purchasePrice,
      "purchaseDate": cryptoData.purchaseDate || 'N/A',
      "totalAmount": cryptoData.totalAmount || (parseInt(cryptoData.quantity) * parseInt(cryptoData.purchasePrice))
    }

    this.dashboardService.saveCryptoInfo(param).subscribe(res => {
      console.log("success")
    }) 
  }

}
