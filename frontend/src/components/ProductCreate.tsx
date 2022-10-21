import React, { useEffect, useState } from "react";
import { Link as RouterLink } from "react-router-dom";
import Button from "@mui/material/Button";
import FormControl from "@mui/material/FormControl";
import Container from "@mui/material/Container";
import Paper from "@mui/material/Paper";
import Grid from "@mui/material/Grid";
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";
import Divider from "@mui/material/Divider";
import Snackbar from "@mui/material/Snackbar";
import Select, { SelectChangeEvent } from "@mui/material/Select";
import MuiAlert, { AlertProps } from "@mui/material/Alert";
import TextField from "@mui/material/TextField";
import { AdapterDateFns } from "@mui/x-date-pickers/AdapterDateFns";
import { LocalizationProvider } from "@mui/x-date-pickers/LocalizationProvider";


import { TypeProductInterface } from "../interfaces/ITypeproduct";
import { ManufacturerInterface } from "../interfaces/IManufacturer";
import { EmployeeInterface } from "../interfaces/IEmployee";
import { ProductInterface } from "../interfaces/IProduct";

import {
  GetTypeproducts,
  GetManufacturer,
  GetEmployees,
  //GetPlaylistByUID,
  Products,
  CreateEmployee,
} from "../services/HttpClientService";
const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(
  props,
  ref
) {
  return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
});

function ProductCreate() {
  const [typeproducts, setTypeproducts] = useState<TypeProductInterface[]>([]);
  const [manufacturers, setManufacturers] = useState<ManufacturerInterface[]>([]);
  const [employees, setEmployees] = useState<EmployeeInterface>();
  const [product, setProduct] = useState<ProductInterface>({});

  const [success, setSuccess] = useState(false);
  const [error, setError] = useState(false);

  const handleInputChange = (
    event: React.ChangeEvent<{ id?: string; value: any }>
  ) => {
    const id = event.target.id as keyof typeof ProductCreate;
    const { value } = event.target;
    setProduct({ ...product, [id]: value });
  };

  
  
  const handleClose = (
    event?: React.SyntheticEvent | Event,
    reason?: string
  ) => {
    if (reason === "clickaway") {
      return;
    }
    setSuccess(false);
    setError(false);
  };
  
  const handleChange = (event: SelectChangeEvent) => {
    const name = event.target.name as keyof typeof product;
    setProduct({
      ...product,
      [name]: event.target.value,
    });
  };

  const getTypeproducts = async () => {
    let res = await GetTypeproducts();
    if (res) {
      setTypeproducts(res);
    }
  };

  const getManufacturer = async () => {
    let res = await GetManufacturer();
    if (res) {
      setManufacturers(res);
    }
  };

  const getEmployee = async () => {
    let res = await GetEmployees();
    if (res) {
      setEmployees(res);
    }
  };

  useEffect(() => {
    getTypeproducts();
    getManufacturer();
    getEmployee();
  }, []);

  const convertType = (data: string | number | undefined) => {
    let val = typeof data === "string" ? parseInt(data) : data;
    return val;
  };



  async function submit() {
    let data = {
      Product_name: product.Product_name??  "",
      Product_price: typeof product.Product_price === "string" ? parseFloat(product.Product_price) : 0,
      TypeproductID: convertType(product.TypeproductID),
      ManufacturerID: convertType(product.ManufacturerID),
      EmployeeID: convertType(product.EmployeeID),
    };

    let res = await Products(data);
    if (res) {
      setSuccess(true);
    } else {
      setError(true);
    }
  }

  return (
    <Container maxWidth="md">
      <Snackbar
        open={success}
        autoHideDuration={3000}
        onClose={handleClose}
        anchorOrigin={{ vertical: "top", horizontal: "center" }}
      >
        <Alert onClose={handleClose} severity="success">
            success
        </Alert>
      </Snackbar>
      <Snackbar
        open={error}
        autoHideDuration={6000}
        onClose={handleClose}
        anchorOrigin={{ vertical: "top", horizontal: "center" }}
      >
        <Alert onClose={handleClose} severity="error">
        Failed to save data
        </Alert>
      </Snackbar>
      <Paper>
        <Box
          display="flex"
          sx={{
            marginTop: 2,
          }}
        >
          <Box sx={{ paddingX: 2, paddingY: 1 }}>
            <Typography
              component="h2"
              variant="h6"
              color="primary"
              gutterBottom
            >
              SAVE PRODUCT INFORMATION
            </Typography>
          </Box>
        </Box>
        <Divider />
        <Grid container spacing={3} sx={{ padding: 2 }}>
          <Grid item xs={12}>
            <p>Name</p>
            <FormControl fullWidth variant="outlined">
              <TextField
                id="Product_name"
                variant="outlined"
                type="string"
                size="medium"
                placeholder="Please select a product name."
                value={product.Product_name || ""}
                onChange={handleInputChange}
              />
            </FormControl>
          </Grid>
          <Grid item xs={12}>
            <FormControl fullWidth variant="outlined">
              <p>Price</p>
              <TextField
                id="Product_price"
                variant="outlined"
                type="number"
                size="medium"
                placeholder="Please select a product price."
                value={product.Product_price || ""}
                onChange={handleInputChange}
              />
            </FormControl>
          </Grid>
          <Grid item xs={12}>
            <FormControl fullWidth variant="outlined">
              <p>Type</p>
              <Select
                native
                value={product.TypeproductID + ""}
                onChange={handleChange}
                inputProps={{
                  name: "TypeproductID",
                }}
              >
                <option aria-label="None" value="">
                    Please select a product type.
                </option>
                {typeproducts.map((item: TypeProductInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Typeproduct_name}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>
          <Grid item xs={12}>
            <FormControl fullWidth variant="outlined">
              <p>Manufacturer</p>
              <Select
                native
                value={product.ManufacturerID + ""}
                onChange={handleChange}
                inputProps={{
                  name: "ManufacturerID",
                }}
              >
                <option aria-label="None" value="">
                Please select a product manufacturer.
                </option>
                {manufacturers.map((item: ManufacturerInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Manufacturer_name}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>
          {/* <Grid item xs={12}>
            <FormControl fullWidth variant="outlined">
              <p>Employee</p>
              <Select
                native
                value={product.EmployeeID + ""}
                onChange={handleChange}
                disabled
                inputProps={{
                  name: "EmployeeID",
                }}
              >
                <option aria-label="None" value="">
                Please select a employee.
                </option>
                <option value={employees?.ID} key={employees?.ID}>
                  {employees?.FirstName}
                </option>

              </Select>
            </FormControl>
          </Grid> */}
          <Grid item xs={12}>
            <p>Employee</p>
            <FormControl fullWidth variant="outlined">
              <TextField
                id="EmployeeID"
                variant="outlined"
                type="number"
                size="medium"
                placeholder="Please select a Employee."
                value={product.EmployeeID || ""}
                onChange={handleInputChange}
              />
            </FormControl>
          </Grid>
          
          <Grid item xs={12}>
            <Button
              component={RouterLink}
              to="/products"
              variant="contained"
              color="inherit"
            >
              BACK
            </Button>
            <Button
              style={{ float: "right" }}
              onClick={submit}
              variant="contained"
              color="primary"
            >
              SAVE
            </Button>
          </Grid>
        </Grid>
      </Paper>
    </Container>
  );
}

export default ProductCreate;