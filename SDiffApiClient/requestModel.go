package SDiffApiClient

type StableDiffusionProcessingTxt2Img struct {
	EnableHr                          bool         `json:"enable_hr,omitempty"`
	DenoisingStrength                 float64      `json:"denoising_strength,omitempty"`
	FirstphaseWidth                   int32        `json:"firstphase_width,omitempty"`
	FirstphaseHeight                  int32        `json:"firstphase_height,omitempty"`
	HrScale                           float64      `json:"hr_scale,omitempty"`
	HrUpscaler                        string       `json:"hr_upscaler,omitempty"`
	HrSecondPassSteps                 int32        `json:"hr_second_pass_steps,omitempty"`
	HrResizeX                         int32        `json:"hr_resize_x,omitempty"`
	HrResizeY                         int32        `json:"hr_resize_y,omitempty"`
	HrSamplerName                     string       `json:"hr_sampler_name,omitempty"`
	HrPrompt                          string       `json:"hr_prompt,omitempty"`
	HrNegativePrompt                  string       `json:"hr_negative_prompt,omitempty"`
	Prompt                            string       `json:"prompt,omitempty"`
	Styles                            []string     `json:"styles,omitempty"`
	Seed                              int32        `json:"seed,omitempty"`
	Subseed                           int32        `json:"subseed,omitempty"`
	SubseedStrength                   float64      `json:"subseed_strength,omitempty"`
	SeedResizeFromH                   int32        `json:"seed_resize_from_h,omitempty"`
	SeedResizeFromW                   int32        `json:"seed_resize_from_w,omitempty"`
	SamplerName                       string       `json:"sampler_name,omitempty"`
	BatchSize                         int32        `json:"batch_size,omitempty"`
	NIter                             int32        `json:"n_iter,omitempty"`
	Steps                             int32        `json:"steps,omitempty"`
	CfgScale                          float64      `json:"cfg_scale,omitempty"`
	Width                             int32        `json:"width,omitempty"`
	Height                            int32        `json:"height,omitempty"`
	RestoreFaces                      bool         `json:"restore_faces,omitempty"`
	Tiling                            bool         `json:"tiling,omitempty"`
	DoNotSaveSamples                  bool         `json:"do_not_save_samples,omitempty"`
	DoNotSaveGrid                     bool         `json:"do_not_save_grid,omitempty"`
	NegativePrompt                    string       `json:"negative_prompt,omitempty"`
	Eta                               float64      `json:"eta,omitempty"`
	SMinUncond                        float64      `json:"s_min_uncond,omitempty"`
	SChurn                            float64      `json:"s_churn,omitempty"`
	STmax                             float64      `json:"s_tmax,omitempty"`
	STmin                             float64      `json:"s_tmin,omitempty"`
	SNoise                            float64      `json:"s_noise,omitempty"`
	OverrideSettings                  *interface{} `json:"override_settings,omitempty"`
	OverrideSettingsRestoreAfterwards bool         `json:"override_settings_restore_afterwards,omitempty"`
	ScriptArgs                        []string     `json:"script_args,omitempty"`
	SamplerIndex                      string       `json:"sampler_index,omitempty"`
	ScriptName                        string       `json:"script_name,omitempty"`
	SendImages                        bool         `json:"send_images,omitempty"`
	SaveImages                        bool         `json:"save_images,omitempty"`
	AlwaysonScripts                   *interface{} `json:"alwayson_scripts,omitempty"`
}