import React, { useState, useEffect } from "react";
import { Link as RouterLink } from "react-router-dom";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import Container from "@mui/material/Container";
import Box from "@mui/material/Box";
import { EmployeeInterface } from "../interfaces/IEmployee";
import { DataGrid, GridColDef } from "@mui/x-data-grid";
import { GetEmployees } from "../services/HttpClientService";

function Employees() {
  const [employees, setEmployees] = useState<EmployeeInterface[]>([]);

  const getEmployes = async () => {
    let res = await GetEmployees();
    if (res) {
      setEmployees(res);
    }
  };

  const columns: GridColDef[] = [
    { field: "ID", headerName: "ลำดับ", width: 50 },
    { field: "Name", headerName: "ชื่อ-สกุล", width: 250,  valueFormatter: (params) => params.value.Name,},
    { field: "Email", headerName: "อีเมลล์", width: 250 , valueFormatter: (params) => params.value.Email,},  
    { field: "Gender", headerName: "เพศ", width: 100 , valueFormatter: (params) => params.value.Name,},    
    { field: "Job_Position", headerName: "ตำแหน่งงาน", width: 200 , valueFormatter: (params) => params.value.Name,}, 
    { field: "Province", headerName: "จังหวัด", width: 200  ,valueFormatter: (params) => params.value.Name,},    
       
  ];

  useEffect(() => {
    getEmployes();
  }, []);

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
              ข้อมูลพนักงาน
            </Typography>
          </Box>
          <Box>
            <Button
              component={RouterLink}
              to="/employee/create"
              variant="contained"
              color="primary"
            >
              สร้างข้อมูล
            </Button>
          </Box>
        </Box>
        <div style={{ height: 400, width: "100%", marginTop: "20px" }}>
          <DataGrid
            rows={employees}
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

export default Employees;