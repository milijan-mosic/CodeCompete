import { useState } from "react";

export const Tasks = () => {
  const [count, setCount] = useState(0);

  return (
    <>
      <h1>Vite + React</h1>

      <div>
        <button onClick={() => setCount((count) => count + 1)}>
          count is {count}
        </button>
      </div>
    </>
  );
};
