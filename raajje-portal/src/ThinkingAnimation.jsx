import React from 'react';

const ThinkingAnimation = ({ message = "Oivaru Processing" }) => {
  return (
    <div className="thinking-container">
      <div className="thinking-core">
        <div className="thinking-ring"></div>
        <div className="thinking-ring" style={{ animationDelay: '0.5s' }}></div>
        <div className="thinking-ring" style={{ animationDelay: '1s' }}></div>
      </div>
      <span className="thinking-text">{message}</span>
    </div>
  );
};

export default ThinkingAnimation;
