import React from 'react';
import PropTypes from 'prop-types';


const border = "border-t-8 border-accent border-solid";


export const StatsCard = ({ children, header, count, measurement, classes="", ...props }) => {
	return (
		<div className={`bg-primary-600 m-1 px-2 pb-1 pt-2 rounded-md ${border} ${classes}`} {...props} >
			<h2>{header}</h2>
			<span className="text-2xl text-accent">{count}</span>
			<span className="font-semibold text-xl tracking-wider"> {measurement}</span>
		</div>
	);
};

StatsCard.propTypes = {
	header: PropTypes.string,
	count: PropTypes.string,
	measurement: PropTypes.string,
	classes: PropTypes.string,
};
