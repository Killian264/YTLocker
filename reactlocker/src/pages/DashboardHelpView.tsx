import { UsingYTLockerInfo, BeforeLinkingInfo, BaseAccountInfo, NeedHelpInfo } from "../components/InfoCards";

export const DashboardHelpView: React.FC<{}> = () => {
	return (
		<div className="w-100">
			<div className="flex flex-col gap-8 w-100 lg:flex-row">
				<div className="flex-grow flex flex-col gap-8 lg:w-2/3">
					<UsingYTLockerInfo></UsingYTLockerInfo>
					<BeforeLinkingInfo></BeforeLinkingInfo>
				</div>
				<div className="flex flex-col gap-8 lg:w-1/3">
					<NeedHelpInfo></NeedHelpInfo>
					<BaseAccountInfo></BaseAccountInfo>
				</div>
			</div>
		</div>
	);
};