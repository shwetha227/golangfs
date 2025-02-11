  
import Greetings from "./Greetings";
import Navbar from "./header/Navbar";

export default function App() {
  return (
    <>
    <Navbar/>
      <h1>
        <marquee>Welcome to NIE</marquee>
      </h1>
      <Greetings/>
    </>
  );
}
