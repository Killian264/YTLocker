import React from 'react';
import PropTypes from 'prop-types';

export const Card = ({ children, className, ...props }) => {
	return (
		<div className={`${className} bg-secondary py-6 px-10 rounded-md`} {...props}>
			{children[0]}
			<div className="border-b-2 mt-2 mb-3"></div>
			{children.slice(1, children.length)}
		</div>
	);
};

Card.propTypes = {
	className: PropTypes.string,
};