import { TypeProductInterface } from "./ITypeproduct";
import { ManufacturerInterface } from "./IManufacturer";
import { EmployeeInterface } from "./IEmployee";

export interface ProductInterface {
  ID?: number;
  Product_name?: string;
  Product_price?: number;
  TypeproductID?: number;
  Typeproduct?: TypeProductInterface;
  ManufacturerID?: number;
  Manufacturer?: ManufacturerInterface;
  EmployeeID?: number;
  Employee?: EmployeeInterface;
}