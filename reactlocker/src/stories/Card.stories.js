import React from 'react';
import { Badge } from '../components/Badge';
import { Card } from '../components/Card';
import { StatsCard } from '../components/StatsCard';
import { toEnum, toBoolean} from "./utils/utils"

export default {
  title: 'Card',
  component: Card,
  argTypes: { onClick: { action: "clicked" } },
};

const Mocked = ({
  children,
  ...props
}) => {
	return (
		<Card {...props} >
			<div className="flex justify-between">
				<div>
					<span className=" text-2xl inline-block align-bottom leading-none">Killian</span>
					<Badge className="ml-1 mt-3" >PRO</Badge>
				</div>
				<div className="flex flex-col justify-end">
					<span className="text-sm leading-none">Joined Mar 28 2020</span>
				</div>
			</div>
			<div className="grid xs:grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 max-w-7xl">
				<StatsCard header={props.header || "Playlists"} count={props.count || "454"} measurement={props.measurement || "total"} />
				<StatsCard header="Videos" count="357" measurement="total" />
				<StatsCard header="Subscriptions" count="17" measurement="total" />
				<StatsCard header="Updated" count="13" measurement="seconds ago" />
			</div>
		</Card>
	);
};


export const Primary = Mocked.bind({})

Primary.argTypes = {};