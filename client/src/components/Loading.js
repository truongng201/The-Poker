import React from "react";
import Spinner from "react-bootstrap/Spinner";

export default function LoadingComponent({ size }) {
  return <Spinner size={size} animation="border" />;
}
