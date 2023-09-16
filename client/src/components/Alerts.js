import "./Alert.css";
import Alert from "react-bootstrap/Alert";

export default function AlertComponent({ message, handleClose }) {
  return (
    <div className="AlertComponent">
      <Alert key="danger" variant="danger" onClose={handleClose} dismissible>
        {message}
      </Alert>
    </div>
  );
}
