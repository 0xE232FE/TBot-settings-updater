package main

type InstanceSetting struct {
	Credentials struct {
		Universe      string `json:"Universe"`
		Email         string `json:"Email"`
		Password      string `json:"Password"`
		Language      string `json:"Language"`
		LobbyPioneers bool   `json:"LobbyPioneers"`
		BasicAuth     struct {
			Username string `json:"Username"`
			Password string `json:"Password"`
		} `json:"BasicAuth"`
		DeviceConf struct {
			Name        string `json:"Name"`
			System      string `json:"System"`
			Browser     string `json:"Browser"`
			UserAgent   string `json:"UserAgent"`
			Memory      int    `json:"Memory"`
			Concurrency int    `json:"Concurrency"`
			Color       int    `json:"Color"`
			Width       int    `json:"Width"`
			Height      int    `json:"Height"`
			Timezone    string `json:"Timezone"`
			Lang        string `json:"Lang"`
		} `json:"DeviceConf"`
	} `json:"Credentials"`
	General struct {
		Host  string `json:"Host"`
		Port  string `json:"Port"`
		Proxy struct {
			Enabled   bool   `json:"Enabled"`
			Address   string `json:"Address"`
			Type      string `json:"Type"`
			Username  string `json:"Username"`
			Password  string `json:"Password"`
			LoginOnly bool   `json:"LoginOnly"`
		} `json:"Proxy"`
		CaptchaAPIKey    string `json:"CaptchaAPIKey"`
		CustomTitle      string `json:"CustomTitle"`
		SlotsToLeaveFree int    `json:"SlotsToLeaveFree"`
	} `json:"General"`
	SleepMode struct {
		Active                  bool   `json:"Active"`
		GoToSleep               string `json:"GoToSleep"`
		WakeUp                  string `json:"WakeUp"`
		PreventIfThereAreFleets bool   `json:"PreventIfThereAreFleets"`
		TelegramMessenger       struct {
			Active bool `json:"Active"`
		} `json:"TelegramMessenger"`
		AutoFleetSave struct {
			Active         bool   `json:"Active"`
			OnlyMoons      bool   `json:"OnlyMoons"`
			DeutToLeave    int    `json:"DeutToLeave"`
			Recall         bool   `json:"Recall"`
			DefaultMission string `json:"DefaultMission"`
		} `json:"AutoFleetSave"`
	} `json:"SleepMode"`
	Defender struct {
		Active              bool `json:"Active"`
		CheckIntervalMin    int  `json:"CheckIntervalMin"`
		CheckIntervalMax    int  `json:"CheckIntervalMax"`
		IgnoreProbes        bool `json:"IgnoreProbes"`
		DefendFromMissiles  bool `json:"DefendFromMissiles"`
		IgnoreWeakAttack    bool `json:"IgnoreWeakAttack"`
		WeakAttackRatio     int  `json:"WeakAttackRatio"`
		IgnoreAttackIfIHave struct {
			Active             bool `json:"Active"`
			MinResourcesToSave int  `json:"MinResourcesToSave"`
			MinFleetToSave     int  `json:"MinFleetToSave"`
		} `json:"IgnoreAttackIfIHave"`
		RandomActivity bool `json:"RandomActivity"`
		Home           struct {
			Galaxy   int    `json:"Galaxy"`
			System   int    `json:"System"`
			Position int    `json:"Position"`
			Type     string `json:"Type"`
		} `json:"Home"`
		Autofleet struct {
			Active            bool `json:"Active"`
			TelegramMessenger struct {
				Active bool `json:"Active"`
			} `json:"TelegramMessenger"`
		} `json:"Autofleet"`
		WhiteList   []int `json:"WhiteList"`
		SpyAttacker struct {
			Active bool `json:"Active"`
			Probes int  `json:"Probes"`
		} `json:"SpyAttacker"`
		MessageAttacker struct {
			Active   bool     `json:"Active"`
			Messages []string `json:"Messages"`
		} `json:"MessageAttacker"`
		TelegramMessenger struct {
			Active bool `json:"Active"`
		} `json:"TelegramMessenger"`
		Alarm struct {
			Active bool `json:"Active"`
		} `json:"Alarm"`
	} `json:"Defender"`
	Brain struct {
		Active            bool `json:"Active"`
		SlotPriorityLevel int  `json:"SlotPriorityLevel"`
		Transports        struct {
			Active                  bool   `json:"Active"`
			CargoType               string `json:"CargoType"`
			DeutToLeave             int    `json:"DeutToLeave"`
			RoundResources          bool   `json:"RoundResources"`
			SendToTheMoonIfPossible bool   `json:"SendToTheMoonIfPossible"`
			Origin                  struct {
				Galaxy   int    `json:"Galaxy"`
				System   int    `json:"System"`
				Position int    `json:"Position"`
				Type     string `json:"Type"`
			} `json:"Origin"`
			MaxSlots                                          int  `json:"MaxSlots"`
			CheckMoonOrPlanetFirst                            bool `json:"CheckMoonOrPlanetFirst"`
			DoMultipleTransportIsNotEnoughShipButSamePosition bool `json:"DoMultipleTransportIsNotEnoughShipButSamePosition"`
			MultipleOrigins                                   struct {
				Active                          bool `json:"Active"`
				OnlyFromMoons                   bool `json:"OnlyFromMoons"`
				MinimumResourcesToSend          int  `json:"MinimumResourcesToSend"`
				PriorityToProximityOverQuantity bool `json:"PriorityToProximityOverQuantity"`
				Exclude                         []struct {
					Galaxy   int    `json:"Galaxy"`
					System   int    `json:"System"`
					Position int    `json:"Position"`
					Type     string `json:"Type"`
				} `json:"Exclude"`
			} `json:"MultipleOrigins"`
		} `json:"Transports"`
		AutoMine struct {
			Active     bool `json:"Active"`
			Transports struct {
				Active bool `json:"Active"`
			} `json:"Transports"`
			MaxMetalMine            int  `json:"MaxMetalMine"`
			MaxCrystalMine          int  `json:"MaxCrystalMine"`
			MaxDeuteriumSynthetizer int  `json:"MaxDeuteriumSynthetizer"`
			MaxSolarPlant           int  `json:"MaxSolarPlant"`
			MaxFusionReactor        int  `json:"MaxFusionReactor"`
			MaxMetalStorage         int  `json:"MaxMetalStorage"`
			MaxCrystalStorage       int  `json:"MaxCrystalStorage"`
			MaxDeuteriumTank        int  `json:"MaxDeuteriumTank"`
			MaxRoboticsFactory      int  `json:"MaxRoboticsFactory"`
			MaxShipyard             int  `json:"MaxShipyard"`
			MaxResearchLab          int  `json:"MaxResearchLab"`
			MaxMissileSilo          int  `json:"MaxMissileSilo"`
			MaxNaniteFactory        int  `json:"MaxNaniteFactory"`
			MaxTerraformer          int  `json:"MaxTerraformer"`
			MaxSpaceDock            int  `json:"MaxSpaceDock"`
			MaxLunarBase            int  `json:"MaxLunarBase"`
			MaxLunarShipyard        int  `json:"MaxLunarShipyard"`
			MaxLunarRoboticsFactory int  `json:"MaxLunarRoboticsFactory"`
			MaxSensorPhalanx        int  `json:"MaxSensorPhalanx"`
			MaxJumpGate             int  `json:"MaxJumpGate"`
			RandomOrder             bool `json:"RandomOrder"`
			Exclude                 []struct {
				Galaxy   int    `json:"Galaxy"`
				System   int    `json:"System"`
				Position int    `json:"Position"`
				Type     string `json:"Type"`
			} `json:"Exclude"`
			OptimizeForStart           bool `json:"OptimizeForStart"`
			PrioritizeRobotsAndNanites bool `json:"PrioritizeRobotsAndNanites"`
			BuildDepositIfFull         bool `json:"BuildDepositIfFull"`
			BuildSolarSatellites       bool `json:"BuildSolarSatellites"`
			BuildCrawlers              bool `json:"BuildCrawlers"`
			DepositHours               int  `json:"DepositHours"`
			MaxDaysOfInvestmentReturn  int  `json:"MaxDaysOfInvestmentReturn"`
			DeutToLeaveOnMoons         int  `json:"DeutToLeaveOnMoons"`
			CheckIntervalMin           int  `json:"CheckIntervalMin"`
			CheckIntervalMax           int  `json:"CheckIntervalMax"`
		} `json:"AutoMine"`
		LifeformAutoMine struct {
			Active     bool `json:"Active"`
			Transports struct {
				Active bool `json:"Active"`
			} `json:"Transports"`
			StartFromCrystalMineLvl   int `json:"StartFromCrystalMineLvl"`
			MaxBasePopulationBuilding int `json:"MaxBasePopulationBuilding"`
			MaxBaseFoodBuilding       int `json:"MaxBaseFoodBuilding"`
			MaxBaseTechBuilding       int `json:"MaxBaseTechBuilding"`
			MaxT2Building             int `json:"MaxT2Building"`
			MaxT3Building             int `json:"MaxT3Building"`
			MaxBuilding6              int `json:"MaxBuilding6"`
			MaxBuilding7              int `json:"MaxBuilding7"`
			MaxBuilding8              int `json:"MaxBuilding8"`
			MaxBuilding9              int `json:"MaxBuilding9"`
			MaxBuilding10             int `json:"MaxBuilding10"`
			MaxBuilding11             int `json:"MaxBuilding11"`
			MaxBuilding12             int `json:"MaxBuilding12"`
			Exclude                   []struct {
				Galaxy   int    `json:"Galaxy"`
				System   int    `json:"System"`
				Position int    `json:"Position"`
				Type     string `json:"Type"`
			} `json:"Exclude"`
			PreventIfMoreExpensiveThanNextMine bool `json:"PreventIfMoreExpensiveThanNextMine"`
			CheckIntervalMin                   int  `json:"CheckIntervalMin"`
			CheckIntervalMax                   int  `json:"CheckIntervalMax"`
		} `json:"LifeformAutoMine"`
		LifeformAutoResearch struct {
			Active     bool `json:"Active"`
			Transports struct {
				Active bool `json:"Active"`
			} `json:"Transports"`
			MaxResearchLevel int `json:"MaxResearchLevel"`
			MaxTechs11       int `json:"MaxTechs11"`
			MaxTechs12       int `json:"MaxTechs12"`
			MaxTechs13       int `json:"MaxTechs13"`
			MaxTechs14       int `json:"MaxTechs14"`
			MaxTechs15       int `json:"MaxTechs15"`
			MaxTechs16       int `json:"MaxTechs16"`
			MaxTechs21       int `json:"MaxTechs21"`
			MaxTechs22       int `json:"MaxTechs22"`
			MaxTechs23       int `json:"MaxTechs23"`
			MaxTechs24       int `json:"MaxTechs24"`
			MaxTechs25       int `json:"MaxTechs25"`
			MaxTechs26       int `json:"MaxTechs26"`
			MaxTechs31       int `json:"MaxTechs31"`
			MaxTechs32       int `json:"MaxTechs32"`
			MaxTechs33       int `json:"MaxTechs33"`
			MaxTechs34       int `json:"MaxTechs34"`
			MaxTechs35       int `json:"MaxTechs35"`
			MaxTechs36       int `json:"MaxTechs36"`
			Exclude          []struct {
				Galaxy   int    `json:"Galaxy"`
				System   int    `json:"System"`
				Position int    `json:"Position"`
				Type     string `json:"Type"`
			} `json:"Exclude"`
			CheckIntervalMin int `json:"CheckIntervalMin"`
			CheckIntervalMax int `json:"CheckIntervalMax"`
		} `json:"LifeformAutoResearch"`
		AutoResearch struct {
			Active     bool `json:"Active"`
			Transports struct {
				Active bool `json:"Active"`
			} `json:"Transports"`
			MaxEnergyTechnology             int `json:"MaxEnergyTechnology"`
			MaxLaserTechnology              int `json:"MaxLaserTechnology"`
			MaxIonTechnology                int `json:"MaxIonTechnology"`
			MaxHyperspaceTechnology         int `json:"MaxHyperspaceTechnology"`
			MaxPlasmaTechnology             int `json:"MaxPlasmaTechnology"`
			MaxCombustionDrive              int `json:"MaxCombustionDrive"`
			MaxImpulseDrive                 int `json:"MaxImpulseDrive"`
			MaxHyperspaceDrive              int `json:"MaxHyperspaceDrive"`
			MaxEspionageTechnology          int `json:"MaxEspionageTechnology"`
			MaxComputerTechnology           int `json:"MaxComputerTechnology"`
			MaxAstrophysics                 int `json:"MaxAstrophysics"`
			MaxIntergalacticResearchNetwork int `json:"MaxIntergalacticResearchNetwork"`
			MaxWeaponsTechnology            int `json:"MaxWeaponsTechnology"`
			MaxShieldingTechnology          int `json:"MaxShieldingTechnology"`
			MaxArmourTechnology             int `json:"MaxArmourTechnology"`
			Target                          struct {
				Galaxy   int    `json:"Galaxy"`
				System   int    `json:"System"`
				Position int    `json:"Position"`
				Type     string `json:"Type"`
			} `json:"Target"`
			OptimizeForStart                       bool `json:"OptimizeForStart"`
			EnsureExpoSlots                        bool `json:"EnsureExpoSlots"`
			PrioritizeAstrophysics                 bool `json:"PrioritizeAstrophysics"`
			PrioritizePlasmaTechnology             bool `json:"PrioritizePlasmaTechnology"`
			PrioritizeEnergyTechnology             bool `json:"PrioritizeEnergyTechnology"`
			PrioritizeIntergalacticResearchNetwork bool `json:"PrioritizeIntergalacticResearchNetwork"`
			CheckIntervalMin                       int  `json:"CheckIntervalMin"`
			CheckIntervalMax                       int  `json:"CheckIntervalMax"`
		} `json:"AutoResearch"`
		AutoCargo struct {
			Active                  bool   `json:"Active"`
			ExcludeMoons            bool   `json:"ExcludeMoons"`
			CargoType               string `json:"CargoType"`
			RandomOrder             bool   `json:"RandomOrder"`
			MaxCargosToBuild        int    `json:"MaxCargosToBuild"`
			MaxCargosToKeep         int    `json:"MaxCargosToKeep"`
			LimitToCapacity         bool   `json:"LimitToCapacity"`
			SkipIfIncomingTransport bool   `json:"SkipIfIncomingTransport"`
			Exclude                 []struct {
				Galaxy   int    `json:"Galaxy"`
				System   int    `json:"System"`
				Position int    `json:"Position"`
				Type     string `json:"Type"`
			} `json:"Exclude"`
			CheckIntervalMin int `json:"CheckIntervalMin"`
			CheckIntervalMax int `json:"CheckIntervalMax"`
		} `json:"AutoCargo"`
		AutoRepatriate struct {
			Active           bool `json:"Active"`
			ExcludeMoons     bool `json:"ExcludeMoons"`
			MinimumResources int  `json:"MinimumResources"`
			LeaveDeut        struct {
				OnlyOnMoons bool `json:"OnlyOnMoons"`
				DeutToLeave int  `json:"DeutToLeave"`
			} `json:"LeaveDeut"`
			Target struct {
				Galaxy   int    `json:"Galaxy"`
				System   int    `json:"System"`
				Position int    `json:"Position"`
				Type     string `json:"Type"`
			} `json:"Target"`
			TargetAssociateMoon     bool   `json:"TargetAssociateMoon"`
			CargoType               string `json:"CargoType"`
			RandomOrder             bool   `json:"RandomOrder"`
			SkipIfIncomingTransport bool   `json:"SkipIfIncomingTransport"`
			Exclude                 []struct {
				Galaxy   int    `json:"Galaxy"`
				System   int    `json:"System"`
				Position int    `json:"Position"`
				Type     string `json:"Type"`
			} `json:"Exclude"`
			CheckIntervalMin int `json:"CheckIntervalMin"`
			CheckIntervalMax int `json:"CheckIntervalMax"`
		} `json:"AutoRepatriate"`
		BuyOfferOfTheDay struct {
			Active           bool `json:"Active"`
			CheckIntervalMin int  `json:"CheckIntervalMin"`
			CheckIntervalMax int  `json:"CheckIntervalMax"`
		} `json:"BuyOfferOfTheDay"`
	} `json:"Brain"`
	Expeditions struct {
		Active                  bool   `json:"Active"`
		SlotPriorityLevel       int    `json:"SlotPriorityLevel"`
		IgnoreSleep             bool   `json:"IgnoreSleep"`
		MinWaitNextFleet        int    `json:"MinWaitNextFleet"`
		MaxWaitNextFleet        int    `json:"MaxWaitNextFleet"`
		PrimaryShip             string `json:"PrimaryShip"`
		MinPrimaryToSend        int    `json:"MinPrimaryToSend"`
		PrimaryToKeep           int    `json:"PrimaryToKeep"`
		SecondaryShip           string `json:"SecondaryShip"`
		MinSecondaryToSend      int    `json:"MinSecondaryToSend"`
		SecondaryToPrimaryRatio int    `json:"SecondaryToPrimaryRatio"`
		ManualShips             struct {
			Active bool `json:"Active"`
			Ships  struct {
				LargeCargo     int `json:"LargeCargo"`
				LightFighter   int `json:"LightFighter"`
				Reaper         int `json:"Reaper"`
				Pathfinder     int `json:"Pathfinder"`
				EspionageProbe int `json:"EspionageProbe"`
			} `json:"Ships"`
		} `json:"ManualShips"`
		WaitForAllExpeditions          bool `json:"WaitForAllExpeditions"`
		WaitForMajorityOfExpeditions   bool `json:"WaitForMajorityOfExpeditions"`
		MinWaitNextRound               int  `json:"MinWaitNextRound"`
		MaxWaitNextRound               int  `json:"MaxWaitNextRound"`
		SplitExpeditionsBetweenSystems struct {
			Active bool `json:"Active"`
			Range  int  `json:"Range"`
		} `json:"SplitExpeditionsBetweenSystems"`
		RandomizeOrder bool `json:"RandomizeOrder"`
		FuelToCarry    int  `json:"FuelToCarry"`
		Origin         []struct {
			Galaxy   int    `json:"Galaxy"`
			System   int    `json:"System"`
			Position int    `json:"Position"`
			Type     string `json:"Type"`
		} `json:"Origin"`
	} `json:"Expeditions"`
	AutoFarm struct {
		Active            bool `json:"Active"`
		SlotPriorityLevel int  `json:"SlotPriorityLevel"`
		ExcludeMoons      bool `json:"ExcludeMoons"`
		ScanRange         []struct {
			Galaxy      int `json:"Galaxy"`
			StartSystem int `json:"StartSystem"`
			EndSystem   int `json:"EndSystem"`
		} `json:"ScanRange"`
		Exclude []struct {
			Galaxy int `json:"Galaxy"`
			System int `json:"System"`
			Planet int `json:"Planet"`
		} `json:"Exclude"`
		KeepReportFor int `json:"KeepReportFor"`
		NumProbes     int `json:"NumProbes"`
		Origin        []struct {
			Galaxy   int    `json:"Galaxy"`
			System   int    `json:"System"`
			Position int    `json:"Position"`
			Type     string `json:"Type"`
		} `json:"Origin"`
		TargetsProbedBeforeAttack int     `json:"TargetsProbedBeforeAttack"`
		CargoType                 string  `json:"CargoType"`
		FleetSpeed                int     `json:"FleetSpeed"`
		MinCargosToKeep           int     `json:"MinCargosToKeep"`
		MinCargosToSend           int     `json:"MinCargosToSend"`
		CargoSurplusPercentage    int     `json:"CargoSurplusPercentage"`
		BuildCargos               bool    `json:"BuildCargos"`
		BuildProbes               bool    `json:"BuildProbes"`
		MinimumResources          int     `json:"MinimumResources"`
		MinimumPlayerRank         int     `json:"MinimumPlayerRank"`
		MaxFlightTime             int     `json:"MaxFlightTime"`
		MaxWaitTime               int     `json:"MaxWaitTime"`
		MinLootFuelRatio          float64 `json:"MinLootFuelRatio"`
		PreferedResource          string  `json:"PreferedResource"`
		MaxSlots                  int     `json:"MaxSlots"`
		SlotsToLeaveFree          int     `json:"SlotsToLeaveFree"`
		CheckIntervalMin          int     `json:"CheckIntervalMin"`
		CheckIntervalMax          int     `json:"CheckIntervalMax"`
	} `json:"AutoFarm"`
	AutoHarvest struct {
		Active                    bool `json:"Active"`
		HarvestOwnDF              bool `json:"HarvestOwnDF"`
		HarvestDeepSpace          bool `json:"HarvestDeepSpace"`
		MinimumResourcesOwnDF     int  `json:"MinimumResourcesOwnDF"`
		MinimumResourcesDeepSpace int  `json:"MinimumResourcesDeepSpace"`
		CheckIntervalMin          int  `json:"CheckIntervalMin"`
		CheckIntervalMax          int  `json:"CheckIntervalMax"`
	} `json:"AutoHarvest"`
	AutoColonize struct {
		Active            bool `json:"Active"`
		SlotPriorityLevel int  `json:"SlotPriorityLevel"`
		Origin            struct {
			Galaxy   int    `json:"Galaxy"`
			System   int    `json:"System"`
			Position int    `json:"Position"`
			Type     string `json:"Type"`
		} `json:"Origin"`
		SlotsToLeaveFree int `json:"SlotsToLeaveFree"`
		Abandon          struct {
			Active                   bool `json:"Active"`
			MinFields                int  `json:"MinFields"`
			MinTemperatureAcceptable int  `json:"MinTemperatureAcceptable"`
			MaxTemperatureAcceptable int  `json:"MaxTemperatureAcceptable"`
		} `json:"Abandon"`
		IntensiveResearch struct {
			Active           bool `json:"Active"`
			MaxSlots         int  `json:"MaxSlots"`
			MinWaitNextFleet int  `json:"MinWaitNextFleet"`
			MaxWaitNextFleet int  `json:"MaxWaitNextFleet"`
		} `json:"IntensiveResearch"`
		Targets []struct {
			Galaxy        int `json:"Galaxy"`
			StartSystem   int `json:"StartSystem"`
			EndSystem     int `json:"EndSystem"`
			StartPosition int `json:"StartPosition"`
			EndPosition   int `json:"EndPosition"`
			MaxPlanets    int `json:"MaxPlanets"`
		} `json:"Targets"`
		RandomPosition   bool `json:"RandomPosition"`
		CheckIntervalMin int  `json:"CheckIntervalMin"`
		CheckIntervalMax int  `json:"CheckIntervalMax"`
	} `json:"AutoColonize"`
	AutoDiscovery struct {
		Active            bool `json:"Active"`
		SlotPriorityLevel int  `json:"SlotPriorityLevel"`
		Origin            struct {
			Galaxy   int    `json:"Galaxy"`
			System   int    `json:"System"`
			Position int    `json:"Position"`
			Type     string `json:"Type"`
		} `json:"Origin"`
		MaxSlots         int `json:"MaxSlots"`
		MaxFailures      int `json:"MaxFailures"`
		CheckIntervalMin int `json:"CheckIntervalMin"`
		CheckIntervalMax int `json:"CheckIntervalMax"`
	} `json:"AutoDiscovery"`
}
