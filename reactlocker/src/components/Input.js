import React from 'react';
import PropTypes from 'prop-types';
import "../styles/tailwind.css"

export const Input = ({ className, ...props }) => {
	const bg = `bg-primary-700`;
	const ring = `focus:outline-none`;
	const c = `w-full py-2 px-4 rounded-md text-primary-100 placeholder-primary-300 focus:text-secondary-text ${ring} ${bg} ${className} `;

	return (
		<input className={c} data-testid="input" {...props} />
	);
};