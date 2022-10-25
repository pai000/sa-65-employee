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
//import { AdapterDateFns } from "@mui/x-date-pickers/AdapterDateFns";
//import { LocalizationProvider } from "@mui/x-date-pickers/LocalizationProvider";
//import { DatePicker } from "@mui/x-date-pickers/DatePicker";

import { ProvincesInterface } from "../interfaces/IProvince"; 
import { GendersInterface } from "../interfaces/IGender";
import { EmployeeInterface } from "../interfaces/IEmployee";

import {
  GetGenders,
  GetJob_Positions,
  GetProvinces,
  CreateEmployees,

} from "../services/HttpClientService";
import { Job_positionsInterface } from "../interfaces/IJob_Position";

const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(
  props,
  ref
) {
  return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
});

function EmployeeCreate() {
  const [gender, setGenders] = useState<GendersInterface[]>([]);
  const [job_position, setJob_Positions] = useState<Job_positionsInterface[]>([]);
  const [provinces, setProvinces] = React.useState<ProvincesInterface[]>([]);
  const [employee, setEmployee] = useState<Partial<EmployeeInterface>>({});

  const [success, setSuccess] = useState(false);
  const [error, setError] = useState(false);

   //เปิดปิดตัว alert
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
    //combobox
    const handleChange = (event: SelectChangeEvent) => {
    const name = event.target.name as keyof typeof EmployeeCreate;
    setEmployee({
      ...employee,
      [name]: event.target.value,
    });
  };

  //text field
  const handleInputChange = (
    event: React.ChangeEvent<{ id?: string; value: any }>
  ) => {
    const id = event.target.id as keyof typeof employee;
    const { value } = event.target;
    setEmployee({ ...employee, [id]: value });
  };
  
  const getGenders = async () => {
    let res = await GetGenders();
    if (res) {
      setGenders(res);
    }
  };

  const getJob_Positions = async () => {
    let res = await GetJob_Positions();
    if (res) {
      setJob_Positions(res);
    }
  };


  const getProvinces = async () => {
    let res = await GetProvinces();
    if (res) {
      setProvinces(res);
    }
  };

  useEffect(() => {
    getGenders();
    getJob_Positions();
    getProvinces();
  }, []);

  const convertType = (data: string | number | undefined) => {
    let val = typeof data === "string" ? parseInt(data) : data;
    return val;
  };


  
  async function submit() {
    let data = {
      Name: employee.Name,
      Email: employee.Email,
      Personal_ID: employee.Personal_ID,
      GenderID: convertType(employee.GenderID),
      Job_PositionID: convertType(employee.Job_PositionID),
      ProvinceID: convertType(employee.ProvinceID),
    };
    console.log(data);

    let res = await CreateEmployees(data);
    console.log(res);
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
          บันทึกข้อมูลสำเร็จ
        </Alert>
      </Snackbar>
      <Snackbar
        open={error}
        autoHideDuration={6000}
        onClose={handleClose}
        anchorOrigin={{ vertical: "top", horizontal: "center" }}
      >
        <Alert onClose={handleClose} severity="error">
          บันทึกข้อมูลไม่สำเร็จ
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
              สร้างข้อมูลพนักงาน
            </Typography>
          </Box>
        </Box>
        <Divider />
        <Grid container spacing={3} sx={{ padding: 2 }}>

        <Grid item xs={6}>
            <p>ชื่อ-สกุล</p>
            <FormControl fullWidth variant="outlined">
              <TextField
                id="Name"
                variant="outlined"
                type="string"
                size="medium"
                placeholder="กรุณากรอกชื่อ-สกุล"
                value={employee.Name || ""}
                onChange={handleInputChange}
              />
            </FormControl>
          </Grid>

          <Grid item xs={6}>
            <p>บัตรประจำตัวประชาชน</p>
            <FormControl fullWidth variant="outlined">
              <TextField
                id="Personal_ID"
                variant="outlined"
                type="string"
                size="medium"
                placeholder="กรุณากรอกบัตรประชาชน"
                value={employee.Personal_ID || ""}
                onChange={handleInputChange}
                inputProps={{maxLength :13}}
              />
            </FormControl>
          </Grid>

          {/* <Grid item xs={6}>
            <p>เลขบัตรประชาชน</p>
            <FormControl fullWidth variant="outlined">
              <TextField
                id="Civ"
                variant="outlined"
                //inputProps={{ inputMode: 'numeric', pattern: '[0-9]*' }}
                type="string"
                size="medium"
                placeholder="กรุณากรอกเลขบัตรประชาชน"
                value={user?.Civ || ""}
                onChange={handleInputChange}
                onKeyPress={(e) => {
                  if (!/[0-9]/.test(e.key)){
                    e.preventDefault()
                  }
                }}
                inputProps={{maxLength :13}}
              />
            </FormControl>
          </Grid> */}

          <Grid item xs={6}>
            <p>อีเมลล์</p>
            <FormControl fullWidth variant="outlined">
              <TextField
                id="Email"
                variant="outlined"
                type="string"
                size="medium"
                placeholder="กรุณากรอกอีเมลล์"
                value={employee.Email || ""}
                onChange={handleInputChange}
              />
            </FormControl>
          </Grid>


          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>เพศ</p>
              <Select
                native
                value={employee.GenderID + ""}
                onChange={handleChange}
                inputProps={{
                  name: "GenderID",
                }}
              >
                <option aria-label="None" value="">
                  กรุณาเลือกเพศ
                </option>
                {gender.map((item: GendersInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Name}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>


          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>ตำแหน่งงาน</p>
              <Select
                native
                value={employee.Job_PositionID + ""}
                onChange={handleChange}
                inputProps={{
                  name: "Job_PositionID",
                }}
              >
                <option aria-label="None" value="">
                  กรุณาเลือกตำแหน่งงาน
                </option>
                {job_position.map((item: Job_positionsInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Name}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>


          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>จังหวัด</p>
              <Select
                native
                value={employee.ProvinceID + ""}
                onChange={handleChange}
                inputProps={{
                  name: "ProvinceID",
                }}
              >
                <option aria-label="None" value="">
                  กรุณาเลือกจังหวัด
                </option>
                {provinces.map((item: ProvincesInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Name}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>

          <Grid item xs={12}>
            <Button
              component={RouterLink}
              to="/employees"
              variant="contained"
              color="inherit"
            >
              กลับ
            </Button>
            <Button
              style={{ float: "right" }}
              onClick={submit}
              variant="contained"
              color="primary"
            >
              บันทึก
            </Button>
          </Grid>
        </Grid>
      </Paper>
    </Container>
  );
}

export default EmployeeCreate;
