import { Paper, Typography } from "@mui/material";
import Dropzone from "react-dropzone/.";

export default function FileZone() {
  return (
    <Dropzone onDrop={(acceptedFiles) => console.log(acceptedFiles)}>
      {({ getRootProps, getInputProps, isDragActive }) => (
        <Paper
          {...getRootProps()}
          variant="outlined"
          sx={{
            width: 500,
            height: 300,
            display: "flex",
            alignItems: "center",
            justifyContent: "center",
            cursor: "pointer",
            borderStyle: "dashed",
            borderColor: isDragActive ? "primary.main" : "divider",
            bgcolor: isDragActive ? "action.hover" : "background.paper",
          }}
        >
          <input {...getInputProps()} />
          <Typography color="text.secondary">
            {isDragActive
              ? "Släpp filerna här"
              : "Dra hit filer eller klicka för att välja"}
          </Typography>
        </Paper>
      )}
    </Dropzone>
  );
}
