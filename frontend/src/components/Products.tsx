import React, { useState, useEffect } from "react";
import { Link as RouterLink } from "react-router-dom";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import Container from "@mui/material/Container";
import Box from "@mui/material/Box";
import { DataGrid, GridColDef } from "@mui/x-data-grid";
import { ProductInterface } from "../interfaces/IProduct";
import { GetProducts } from "../services/HttpClientService";

function Products() {
  const [products, setProducts] = useState<ProductInterface[]>([]);

  useEffect(() => {
    getProducts();
  }, []);

  const getProducts = async () => {
    let res = await GetProducts();
    if (res) {
      setProducts(res);
    } 
  };

  const columns: GridColDef[] = [
    { field: "ID", headerName: "ID", width: 70 },
    {
      field: "Product_name",
      headerName: "Name",
      width: 200,
      valueFormatter: (params) => params.value.Product_name,
    },
    {
        field: "Product_price",
        headerName: "Price",
        width: 100,
        valueFormatter: (params) => params.value.Product_price,
      },
    {
      field: "Typeproduct",
      headerName: "Typeproduct",
      width: 150,
      valueFormatter: (params) => params.value.Typeproduct_name,
    },
    {
      field: "Manufacturer",
      headerName: "Manufacturer",
      width: 150,
      valueFormatter: (params) => params.value.Manufacturer_name,
    },
    { field: "Employee", 
      headerName: "Employee", 
      width: 150 ,
      valueFormatter: (params) => params.value.FirstName + " " + params.value.LastName,
    },

  ];

  return (
    <div>
      <Container maxWidth="md">
        <Box
          display="flex"
          sx={{
            marginTop: 2,
          }}
        >
          <Box flexGrow={1}>
            <Typography
              component="h2"
              variant="h6"
              color="primary"
              gutterBottom
            >
              Product Information
            </Typography>
          </Box>
          <Box>
            <Button
              component={RouterLink}
              to="/product/create"
              variant="contained"
              color="primary"
            >
              Create Data
            </Button>
          </Box>
        </Box>
        <div style={{ height: 400, width: "100%", marginTop: "20px" }}>
          <DataGrid
            rows={products}
            getRowId={(row) => row.ID}
            columns={columns}
            pageSize={5}
            rowsPerPageOptions={[5]}
          />
        </div>
      </Container>
    </div>
  );
}

export default Products;